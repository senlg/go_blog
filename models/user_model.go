package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Role int

func (r *Role) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(r.toString(), "", " ")
}
func (r Role) toString() (str string) {
	switch r {
	case PermissionAdmin:
		str = "管理员"
	case PermissionVisitor:
		str = "游客"
	case PermissionUser:
		str = "普通用户"
	default:
		str = "未匹配用户"
	}
	return
}

const (
	PermissionAdmin   Role = iota + 1 //管理员
	PermissionUser                    //普通用户
	PermissionVisitor                 //游客
)

// Token          string         `gorm:"size:128" json:"token"`                                                            // token
// 用户表
type UserModel struct {
	Model
	NickName       string         `gorm:"size:42" json:"nick_name"`                                                         //昵称
	UserName       string         `gorm:"size:36" json:"user_name"`                                                         //用户名
	Password       string         `gorm:"size:32" json:"password"`                                                          //密码
	AvatarUrl      string         `gorm:"size:256" json:"avatar_url"`                                                       //头像地址
	Addr           string         `gorm:"size:256" json:"addr"`                                                             //地址
	Role           Role           `gorm:"size:12;default:2" json:"role"`                                                    // 用户权限
	Email          string         `json:"email"`                                                                            // 邮箱
	Phone          string         `gorm:"size:30" json:"phone"`                                                             //手机号
	ArticleModels  []ArticleModel `gorm:"foreignKey:UserId" json:"-"`                                                       //发布的文章
	CollectsModels []ArticleModel `gorm:"many2many:user_collects;joinForeignKey:UserId;joinReferences:ArticleId" json:"-" ` //收藏的文章
}

func (UserModel) TableName() string {
	return "user_models"
}

// 登录记录
type LoginRecordModel struct {
	Id          uint      `gorm:"primaryKey"`
	LoginTime   time.Time `json:"login_time"`
	LoginAdress string    `json:"login_adress"`
	IP          string    `json:"ip"`
	UserId      uint      `json:"-"`
	UserModel   UserModel `gorm:"foreignKey:UserId"`
}

// 自定义收藏文章表
type UserCollect struct {
	CreatedAt time.Time
	gorm.DeletedAt
	UserId    uint `gorm:"primaryKey"`
	ArticleId uint `gorm:"primaryKey"`
}
