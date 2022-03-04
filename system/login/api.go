package login

import (
	"errors"
	"fmt"
	"time"

	"github.com/fangjie-luoxi/gommon/random"
	"github.com/fangjie-luoxi/tools/jwt"
	"github.com/fangjie-luoxi/tools/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// OpenApi 登录相关开放api
func (l *Login) OpenApi(r *gin.RouterGroup, jwt *jwt.GinJWTMiddleware) {
	r.POST("/login", jwt.LoginHandler) // 登录
	r.POST("/register", l.Register)    // 注册
	r.GET("/sms", l.GetSmsCode)        // 获取短信
}

// Api 用户相关api，需要token
func (l *Login) Api(r *gin.RouterGroup, jwt *jwt.GinJWTMiddleware) {
	r.GET("/refresh_token", jwt.RefreshHandler) // 刷新token
	r.GET("/logout", jwt.LogoutHandler)         // 注销
	r.GET("/login_user", l.GetUserLoginInfo)    // 获取用户登录信息
}

// GetSmsCode 获取短信验证码
func (l *Login) GetSmsCode(c *gin.Context) {
	phone := c.Query("Phone")
	if phone == "" {
		response.Error(c, 400, "Phone不能为空")
		return
	}
	param := map[string]string{
		"SignName":     l.Config.DefaultString("dysms.SignName", "阿里云短信测试"),
		"TemplateCode": l.Config.DefaultString("dysms.TemplateCode", "SMS_154950909"),
	}
	param["PhoneNumbers"] = phone
	code := random.String(6, random.Numeric)
	l.cache.Set(phone, code, 5*time.Minute)
	codeName := l.Config.DefaultString("dysms.code_name", "code")
	param["TemplateParam"] = fmt.Sprintf(`{"%s":"%s"}`, codeName, code)
	b, err := l.sms.SendSms(param)
	if !b {
		response.Error(c, 500, "获取验证码失败:"+err.Error())
		return
	}
	response.Success(c, 200, "OK")
}

// GetUserLoginInfo 获取用户登录信息
func (l *Login) GetUserLoginInfo(c *gin.Context) {
	user, _ := c.Get(l.IdentityKey)
	res, ok := user.(*User)
	if !ok {
		response.Error(c, 400, "获取token信息失败")
		return
	}
	err := l.model.First(&res).Error
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.Success(c, 200, res)
}

// Register 用户注册
func (l *Login) Register(c *gin.Context) {
	var r RegisterInfo
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	user := User{
		SysOrgId:  0,
		LoginCode: r.LoginCode,
		UserCode:  r.LoginCode,
		UserName:  r.UserName,
		Phone:     r.Phone,
		Email:     r.Email,
	}

	if !errors.Is(l.model.Where("login_code = ?", r.LoginCode).First(&User{}).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		response.Error(c, 400, "用户已存在")
		return
	}
	if !errors.Is(l.model.Where("email = ?", r.Email).First(&User{}).Error, gorm.ErrRecordNotFound) && r.Email != "" {
		response.Error(c, 400, "邮箱已注册")
		return
	}
	if r.Phone != "" && !errors.Is(l.model.Where("phone = ?", r.Phone).First(&User{}).Error, gorm.ErrRecordNotFound) {
		response.Error(c, 400, "手机号已注册")
		return
	}
	// md5加密
	user.Password = encrypt(r.Password)
	err = l.model.Create(&user).Error
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.Success(c, 200, "注册成功")
}
