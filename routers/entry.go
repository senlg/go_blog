package routers

import (
	"fmt"
	"go_blog/global"
	"go_blog/middleware"
	"go_blog/routers/api"
	"go_blog/routers/static"
	"time"

	_ "go_blog/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {

	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	count := 0
	router.GET("sse", func(ctx *gin.Context) {
		// 设置响应头
		ctx.Writer.Header().Set("Content-Type", "text/event-stream")
		ctx.Writer.Header().Set("Cache-Control", "no-cache")
		ctx.Writer.Header().Set("Connection", "keep-alive")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		for {
			count = count + 1
			ctx.SSEvent("message", fmt.Sprintf("hello sse client %d", count))
			ctx.Writer.Flush()
			time.Sleep(5 * time.Second)
		}
	})

	// api文档网址
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
