package login

import (
	"time"

	"gorm.io/gorm"
)

// User 用户表
type User struct {
	Id        int            `gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index:idx_del" json:"-"`
	SysOrgId  int            `gorm:"type:int(11);index:idx_org"`                             // 机构id
	LoginCode string         `gorm:"type:varchar(45);index:idx_login_code" binding:"max=45"` // 登录账号
	Password  string         `gorm:"type:varchar(255)" binding:"max=255" json:"-"`           // 密码
	UserCode  string         `gorm:"type:varchar(45);index:idx_user_code" binding:"max=45"`  // 用户编码
	UserName  string         `gorm:"type:varchar(45)" binding:"max=45"`                      // 用户名
	Phone     string         `gorm:"type:varchar(45);index:idx_phone" binding:"max=45"`      // 电话号码
	Avatar    string         `gorm:"type:varchar(255)" binding:"max=255"`                    // 头像
	Email     string         `gorm:"type:varchar(255);index:idx_email" binding:"max=255"`    // 邮箱
	UserType  string         `gorm:"type:varchar(45)"`                                       // 用户类型 `gorm:"type:varchar(10);default:user;not null;comment:用户类型 enum:super:超级管理员,admin:机构管理员,user:普通用户#" binding:"max=10,required"`
	Status    string         `gorm:"type:varchar(10)"`                                       // 用户状态 default:r;comment:状态 enum:r:正常,c:禁用#
	Integral  float64        `gorm:"column:integral"`
}

func (m *User) TableName() string {
	return "user"
}

// ThirdAuth 第三方授权表
type ThirdAuth struct {
	Id        int            `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index:idx_del" json:"-"`
	AuthType  string         `gorm:"column:auth_type" binding:"max=10"`                  // 认证类型 enum:wechat:微信,work:企业微信,dd:钉钉#
	Openid    string         `gorm:"column:openid;index:idx_openid" binding:"max=128"`   // 第三方openid
	Unionid   string         `gorm:"column:unionid;index:idx_unionid" binding:"max=255"` // 第三方unionid
	Nickname  string         `gorm:"column:nickname" binding:"max=64"`                   // 用户昵称
	Avatar    string         `gorm:"column:avatar" binding:"max=255"`                    // 头像
	UserId    *int           `gorm:"column:user_id;index:idx_user"`
	User      *User          // 一对多,从表
}

func (m *ThirdAuth) TableName() string {
	return "third_auth"
}
