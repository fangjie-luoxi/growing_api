package middleware

import (
	_ "embed"
	"time"

	"github.com/fangjie-luoxi/tools/jwt"
	"github.com/gin-gonic/gin"

	"github.com/fangjie-luoxi/growing_api/global"
	"github.com/fangjie-luoxi/growing_api/system/login"
	"github.com/fangjie-luoxi/growing_api/system/system/model"
)

//go:embed keys/priv.key
var privKey string // 静态文件私钥

//go:embed keys/pub.key
var pubKey string // 静态文件公钥

var timeout = 24 * time.Hour    // token过期时间
var maxRefresh = 24 * time.Hour // 允许刷新时间
var algorithm = "RS256"         // 签名算法 HS256
var privKeyFile = ""            // 私钥文件 keys/priv.key
var pubKeyFile = ""             // 公钥文件 keys/pub.key
var hs256Key = ""               // 当algorithm为HS256时的密钥

func JWT(login *login.Login) (*jwt.GinJWTMiddleware, error) {
	openJwt := login.Config.DefaultBool("openJwt", true)
	return jwt.New(&jwt.GinJWTMiddleware{
		SigningAlgorithm: algorithm,
		PrivKeyFile:      privKeyFile,
		PrivKeyBytes:     []byte(privKey),
		PubKeyFile:       pubKeyFile,
		PubKeyBytes:      []byte(pubKey),
		Key:              []byte(hs256Key),
		Timeout:          timeout,
		MaxRefresh:       maxRefresh,
		IdentityKey:      login.IdentityKey,     // 身份密钥
		PayloadFunc:      login.PayloadFunc,     // 登录期间将调用的回调函数。 使用此功能可以将其他有效负载数据添加到JWT令牌
		IdentityHandler:  login.IdentityHandler, // 设置身份处理程序功能
		Authenticator:    login.Authenticator,   // 认证处理程序
		Authorizator:     login.Authorizator,    // 身份认证
		LoginResponse:    login.Response,        // 登录返回
		// LogoutResponse:   login.LogoutResponse,  // 注销接口
		// 登录失败或身份认证失败回调
		Unauthorized: func(c *gin.Context, code int, message string) {
			if openJwt && code == 401 && isOpenJwt(c) {
				return
			}
			c.Abort()
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		DisabledAbort: openJwt,
		// 用于从请求中提取令牌
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// 通过Cookie授权
		SendCookie:     true,
		CookieHTTPOnly: true,
		// 提供当前时间
		TimeFunc: time.Now,
	})
}

// isOpenJwt 是否是开放的jwt
func isOpenJwt(c *gin.Context) bool {
	method := c.Request.Method
	path := c.Request.URL.Path
	apiCache := getApiJwtCache()
	if mapMatch(apiCache, path, method) {
		return true
	}
	return false
}

func getApiJwtCache() map[string][]string {
	key := "ApiCacheJwt"
	if global.Cache.Exists(key) {
		if d, found := global.Cache.Get(key); found {
			if data, ok := d.(map[string][]string); ok {
				return data
			}
		}
	} else {
		var apis []model.SysApi
		global.DBEngine.Where(map[string]interface{}{"api_tp": "jwt"}).Find(&apis)
		apiMap := make(map[string][]string)
		for _, api := range apis {
			methods := apiMap[api.Path]
			methods = append(methods, api.Method)
			apiMap[api.Path] = methods
		}
		global.Cache.Set(key, apiMap, -1)
		return apiMap
	}
	return map[string][]string{}
}
