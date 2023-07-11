package api

import (
	"go_blog/api"

	"github.com/gin-gonic/gin"
)

func InitImageApi(apiGroup *gin.RouterGroup) {
	var i api.ImageApi
	apiGroup.POST("get_image_list", i.FindImageList)
	apiGroup.POST("change_img_info", i.ChangeImageInfo)
	apiGroup.DELETE("delete_images", i.DeleteImage)
}
