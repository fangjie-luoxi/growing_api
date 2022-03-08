package login

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"log"
	"strconv"
	"time"

	"github.com/fangjie-luoxi/ali-sms"
	"github.com/fangjie-luoxi/tools/config"
	"github.com/fangjie-luoxi/tools/jwt"
	"gorm.io/gorm"

	"github.com/fangjie-luoxi/growing_api/pkg/wechat"
)

// ParamLogin 登录参数
type ParamLogin struct {
	LoginCode string `binding:"required"` // 账号
	Password  string // 密码

	LoginType string            // 登录类型 enum:psw:密码登录,qw:企业微信扫码登录,wc:微信扫码登录,mini:小程序扫码登录# 默认psw
	Param     map[string]string // 附加参数
}

type Login struct {
	model  *gorm.DB       // go驱动
	cache  *cache.Cache   // 缓存
	wechat *wechat.Wechat // 微信
	sms    SmsServer      // 短信服务

	Config      *config.Config // 配置
	IdentityKey string         // 身份密钥
}

// SmsServer 短信服务接口
type SmsServer interface {
	SendSms(param map[string]string) (bool, error)
}

func NewLogin(m *gorm.DB, conf *config.Config, wechat *wechat.Wechat) *Login {
	// 初始化数据库
	err := m.AutoMigrate(&User{}, &ThirdAuth{})
	if err != nil {
		log.Println(err)
	}
	accessKey := conf.String("dysms.access_key")
	accessSecret := conf.String("dysms.access_secret")
	return &Login{
		model:  m,
		cache:  cache.New(10*time.Minute, 5*time.Minute),
		sms:    ali_sms.NewSmsClient(accessKey, accessSecret),
		wechat: wechat,

		Config:      conf,
		IdentityKey: "identity",
	}
}

// 账号密码登录
func (l *Login) accountLogin(param *ParamLogin) (*User, error) {
	var user User
	err := l.model.
		Where("login_code = @code OR user_code = @code OR phone = @code OR email = @code", map[string]interface{}{"code": param.LoginCode}).
		First(&user).Error
	password := encrypt(param.Password)
	if err != nil || user.Password != password {
		return nil, jwt.ErrFailedAuthentication
	}
	return &user, nil
}

// 微信小程序授权登录
func (l *Login) miniLogin(param *ParamLogin) (*User, error) {
	var user User
	var thirdAuth ThirdAuth
	userType := "user"
	phone := param.Param["phone"]

	session, err := l.wechat.MiniProgram.GetAuth().Code2Session(param.LoginCode)
	if err != nil {
		return nil, err
	}

	err = l.model.Preload("User").Where(ThirdAuth{Openid: session.OpenID}).Assign(ThirdAuth{AuthType: "wechat"}).FirstOrCreate(&thirdAuth).Error
	if err != nil {
		return nil, err
	}
	if thirdAuth.User == nil {
		assignUser := User{
			UserCode: strconv.Itoa(int(time.Now().Unix())),
			UserName: "微信用户",
			UserType: userType,
			SysOrgId: 0,
			Integral: 1000,
		}
		if phone != "" {
			err = l.model.Where(User{Phone: phone}).Assign(assignUser).First(&user).Error
		} else {
			err = l.model.Create(&assignUser).Error
			user = assignUser
		}
		if user.Id != 0 {
			thirdAuth.UserId = &user.Id
			err = l.model.Save(&thirdAuth).Error
		}
	} else {
		user = *thirdAuth.User
	}
	return &user, err
}

// 企业微信授权登录
func (l *Login) workLogin(param *ParamLogin) (*User, error) {
	var user User
	var thirdAuth ThirdAuth
	auth := l.wechat.Work.GetOauth()
	userInfo, err := auth.UserFromCode(param.LoginCode)
	if err != nil {
		return nil, err
	}
	if userInfo.UserID != "" {
		err = l.model.Preload("User").First(&thirdAuth, "unionid = ?", userInfo.UserID).Error
	} else {
		err = l.model.Preload("User").First(&thirdAuth, "openid = ?", userInfo.OpenID).Error
	}

	if thirdAuth.User == nil {
		thirdAuth.AuthType = "work"
		thirdAuth.Openid = userInfo.OpenID
		thirdAuth.Unionid = userInfo.UserID
		user = User{
			UserCode: strconv.Itoa(int(time.Now().Unix())),
			UserName: "用户" + strconv.Itoa(int(time.Now().Unix())),
			//UserType: userType,
			SysOrgId: 0,
			//Phone:    phone,
		}
		thirdAuth.User = &user
		l.model.Create(&thirdAuth)
	} else {
		user = *thirdAuth.User
	}
	return &user, err
}

// 微信扫码登录 开放平台登录
func (l *Login) openPlatLogin(param *ParamLogin) (*User, error) {
	var user User
	var thirdAuth ThirdAuth

	oauth := l.wechat.OfficialAccount.GetOauth()
	access, err := oauth.GetUserAccessToken(param.LoginCode)
	if err != nil {
		return nil, err
	}
	userInfo, err := oauth.GetUserInfo(access.AccessToken, access.OpenID, "")

	err = l.model.Preload("User").First(&thirdAuth, "openid = ?", userInfo.OpenID).Error
	if err != nil {
		thirdAuth.AuthType = "work"
		thirdAuth.Openid = userInfo.OpenID
		thirdAuth.Unionid = userInfo.Unionid
		thirdAuth.Nickname = userInfo.Nickname
		thirdAuth.Avatar = userInfo.HeadImgURL
		user = User{
			UserCode: strconv.Itoa(int(time.Now().Unix())),
			UserName: userInfo.Nickname,
			//UserType: userType,
			SysOrgId: 0,
			//Phone:    phone,
		}
		thirdAuth.User = &user
		l.model.Create(&thirdAuth)
	} else {
		user = *thirdAuth.User
	}
	return &user, err
}

// 短信登录
func (l *Login) smsLogin(param *ParamLogin) (*User, error) {
	phone := param.LoginCode
	if d, b := l.cache.Get(param.LoginCode); b {
		s, ok := d.(string)
		if !ok {
			return nil, errors.New("验证码错误")
		}
		if param.Password != s || s == "" {
			return nil, errors.New("验证码错误")
		}
	} else {
		return nil, errors.New("验证码错误")
	}

	var user User
	l.model.Where("phone = ?", phone).First(&user)
	if user.Id == 0 {
		user = User{
			UserCode: strconv.Itoa(int(time.Now().Unix())),
			UserName: "用户" + strconv.Itoa(int(time.Now().Unix())),
			Phone:    phone,
			UserType: "user",
			SysOrgId: 0,
		}
		err := l.model.Create(&user).Error
		if err != nil {
			return nil, err
		}
	}
	l.cache.Delete(param.LoginCode)
	return &user, nil
}
