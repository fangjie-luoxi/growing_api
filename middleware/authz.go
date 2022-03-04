package middleware

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fangjie-luoxi/growing_api/global"
	"github.com/fangjie-luoxi/growing_api/pkg/cache"
	"github.com/fangjie-luoxi/growing_api/system/login"
	"github.com/fangjie-luoxi/growing_api/system/system/model"
)

// NewAuthorizer 授权中间件
func NewAuthorizer(e *casbin.Enforcer, o *gorm.DB, cache *cache.Cache) gin.HandlerFunc {
	a := &BasicAuthorizer{enforcer: e, orm: o, apiCache: cache}
	return func(c *gin.Context) {
		if !a.CheckPermission(c) {
			a.RequirePermission(c)
		}
	}
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
	orm      *gorm.DB
	apiCache *cache.Cache // api信息缓存 键SysOrgId，值SysApi map 键 api.Method+":"+api.Path 值SysApi ["/get"]{Path:....}
}

// GetUserName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func (a *BasicAuthorizer) GetUserName(r *http.Request) string {
	username, _, _ := r.BasicAuth()
	return username
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(c *gin.Context) bool {
	userI, exists := c.Get(global.Login.IdentityKey)
	if !exists {
		return true
	}
	user, ok := userI.(*login.User)
	if !ok {
		return true
	}
	method := c.Request.Method
	path := c.Request.URL.Path

	apiCache := a.getApiCache(user.SysOrgId)
	if mapMatch(apiCache, path, method) {
		enforce, err := a.enforcer.Enforce("user::"+strconv.Itoa(user.Id), path, method)
		if err != nil {
			log.Println("Enforce err:", err)
		}
		return enforce
	}
	return true
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

// getApiCache 获取api缓存
func (a *BasicAuthorizer) getApiCache(sysOrgId int) map[string][]string {
	key := "ApiCachePerm:" + strconv.Itoa(sysOrgId)
	if a.apiCache.Exists(key) {
		if d, found := a.apiCache.Get(key); found {
			if data, ok := d.(map[string][]string); ok {
				return data
			}
		}
	} else {
		var apis []model.SysApi
		a.orm.Where(map[string]interface{}{"sys_org_id": sysOrgId, "api_tp": "perm"}).Find(&apis)
		apiMap := make(map[string][]string)
		for _, api := range apis {
			methods := apiMap[api.Path]
			methods = append(methods, api.Method)
			apiMap[api.Path] = methods
		}
		a.apiCache.Set(key, apiMap, -1)
		return apiMap
	}
	return map[string][]string{}
}

// mapMatch 元素是否与map匹配
func mapMatch(apiCache map[string][]string, path, method string) bool {
	for k, v := range apiCache {
		if keyMatch(path, k) {
			for _, cacheMethod := range v {
				regexMatch, err := regexp.MatchString(cacheMethod, method)
				if err != nil {
					log.Println("错误的正则:", err)
					return false
				}
				if regexMatch {
					return true
				}
			}
		}
	}
	return false
}

// keyMatch determines whether key1 matches the pattern of key2 (similar to RESTful path), key2 can contain a *.
// For example, "/foo/bar" matches "/foo/*"
func keyMatch(key1 string, key2 string) bool {
	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}

	if len(key1) > i {
		return key1[:i] == key2[:i]
	}
	return key1 == key2[:i]
}
