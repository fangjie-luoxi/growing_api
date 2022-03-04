package login

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/fangjie-luoxi/tools/jwt"
	"github.com/gin-gonic/gin"
)

// ClaimsInfo token携带的信息，根据需要自定义
type ClaimsInfo struct {
	UserId   int    // 用户id
	UserName string // 用户名
	UserType string // 用户类型
}

// RegisterInfo 注册参数
type RegisterInfo struct {
	LoginCode string `binding:"max=45,required"`         // 登录账号
	UserName  string `binding:"max=45,required"`         // 用户名
	Password  string `binding:"max=255,required"`        // 密码
	Phone     string `binding:"omitempty,max=45"`        // 手机号码
	Email     string `binding:"omitempty,max=255,email"` // 邮箱
}

// Authenticator 登录处理程序
func (l *Login) Authenticator(c *gin.Context) (interface{}, error) {
	var loginVal ParamLogin
	if err := c.ShouldBindJSON(&loginVal); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	var user *User
	var err error

	switch loginVal.LoginType {
	case "work":
		user, err = l.workLogin(&loginVal)
	case "wc":
		user, err = l.openPlatLogin(&loginVal)
	case "mini":
		user, err = l.miniLogin(&loginVal)
	case "sms":
		user, err = l.smsLogin(&loginVal)
	default:
		user, err = l.accountLogin(&loginVal)
	}
	if err != nil {
		return "", err
	}

	if user.Status == "c" {
		return "", errors.New("用户已被禁用")
	}
	c.Set("user", user)

	return &ClaimsInfo{
		UserId:   user.Id,
		UserName: user.UserName,
		UserType: user.UserType,
	}, nil
}

// Response 登录返回
func (l *Login) Response(c *gin.Context, code int, token string, expire time.Time) {
	res := gin.H{
		"code":   code,
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	}
	value, exists := c.Get("user")
	if exists {
		res["user"] = value
	}
	c.JSON(200, res)
}

// PayloadFunc 设置token密钥
func (l *Login) PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*ClaimsInfo); ok {
		return jwt.MapClaims{
			"UserId":   v.UserId,
			"UserName": v.UserName,
			"UserType": v.UserType,
		}
	}
	return jwt.MapClaims{}
}

// IdentityHandler 身份信息处理
func (l *Login) IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &User{
		Id:       int(claims["UserId"].(float64)),
		UserName: claims["UserName"].(string),
		UserType: claims["UserType"].(string),
	}
}

// Authorizator 身份认证处理
func (l *Login) Authorizator(data interface{}, c *gin.Context) bool {
	return true
}

// 加密程序
func encrypt(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
