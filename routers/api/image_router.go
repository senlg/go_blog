package api

import (
	"go_blog/api"

	"github.com/gin-gonic/gin"
)

func InitImageApi(apiGroup *gin.RouterGroup) {
	i := api.ImageApi{}
	apiGroup.POST("get_image_list", i.FindImageList)
}
