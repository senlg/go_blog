package api

import (
	"go_blog/api"

	"github.com/gin-gonic/gin"
)

func InitUserGroupApi(apiGroup *gin.RouterGroup) {
	var u api.User
	apiGroup.POST("getUser", u.GetUserInfo)
	apiGroup.POST("create_user", u.CreateUserInfo)
	apiGroup.POST("login", u.Login)
}
