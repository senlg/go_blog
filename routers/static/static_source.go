package static

import (
	"go_blog/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitStaticSource(r *gin.Engine) {
	r.StaticFS("/static", http.Dir(global.Config.UploadConfig.BasePath))

	// r.Static("/static", global.Config.UploadConfig.BasePath)
}
