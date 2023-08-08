package api

import (
	"go_blog/api"

	"github.com/gin-gonic/gin"
)

func InitMenuApi(apiGroup *gin.RouterGroup) {
	var m api.MenuApi
	apiGroup.GET("get_menu", m.GetMenuList)
	apiGroup.POST("create_menu", m.CreateMenu)
}
