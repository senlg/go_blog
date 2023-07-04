package routers

import (
	"go_blog/global"
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
	api.InitUserGroupApi(apiGroup)
	// 上传api
	api.InitUploadApi(apiGroup)
	// 图片api
	api.InitImageApi(apiGroup)
	return router
}
