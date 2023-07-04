package api

import (
	"go_blog/api"

	"github.com/gin-gonic/gin"
)

func InitUploadApi(apiGroup *gin.RouterGroup) {
	img := api.UploadFile{
		RequestKey: "images",
	}
	apiGroup.POST("upload_image", img.Upload)
	file := api.UploadFile{
		RequestKey: "files",
	}
	apiGroup.POST("upload_file", file.Upload)
}
