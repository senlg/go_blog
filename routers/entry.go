package routers

import (
	"go_blog/global"
	"go_blog/middleware"
	"go_blog/routers/api"
	"go_blog/routers/static"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	// 静态文件api
	static.InitStaticSource(router)
	apiGroup := router.Group("api")
	// jwt中间件
	apiGroup.Use(middleware.JwtMiddleware())
	// 用户api
	api.InitUserGroupApi(apiGroup)
	// 上传api
	api.InitUploadApi(apiGroup)
	// 图片api
	api.InitImageApi(apiGroup)
	// 标签api
	api.InitTagApi(apiGroup)
	// 文章api
	api.InitArticleApi(apiGroup)
	// 留言api
	api.InitCommoentApi(apiGroup)
	// 菜单api
	api.InitMenuApi(apiGroup)

	return router
}
