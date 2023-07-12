package api

import (
	"go_blog/api"

	"github.com/gin-gonic/gin"
)

func InitUserGroupApi(apiGroup *gin.RouterGroup) {
	var u api.User
	apiGroup.POST("get_user", u.GetUserInfo)
	apiGroup.POST("get_user_list", u.GetUserInfoList)
	apiGroup.POST("create_user", u.CreateUserInfo)
	apiGroup.POST("login", u.Login)
	apiGroup.DELETE("del_user", u.DelUser)
}
