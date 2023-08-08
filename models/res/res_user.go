package res

import "go_blog/models"

type UserInfo struct {
	Id            uint        `gorm:"primarykey" json:"id"`
	UserName      string      `json:"user_name"`
	NickName      string      `json:"nick_name"`
	AvatarUrl     string      `json:"avatar_url"`
	Addr          string      `json:"addr"`
	Token         string      `json:"token"`
	Role          models.Role `json:"role"`
	RoleName      string      `json:"role_name"`
	Phone         string      `json:"phone"`
	ReleaseCount  int         `json:"release_count"`
	CollectsCount int         `json:"collects_count"`
}

type UserItem struct {
	Id           uint        `gorm:"primarykey" json:"id"`
	UserName     string      `json:"user_name"`
	NickName     string      `json:"nick_name"`
	AvatarUrl    string      `json:"avatar_url"`
	Addr         string      `json:"addr"`
	Role         models.Role `json:"role"`
	Phone        string      `json:"phone"`
	ReleaseCount int         `json:"release_count"`
}
