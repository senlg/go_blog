package api

import (
	"go_blog/api"

	"github.com/gin-gonic/gin"
)

func InitTagApi(apiGroup *gin.RouterGroup) {
	var t api.TagApi
	apiGroup.POST("get_tag_list", t.GetTagList)
	apiGroup.POST("create_tag", t.CreateTag)
	apiGroup.POST("update_tag", t.UpdateTag)
	apiGroup.DELETE("delete_tags", t.DeleteTag)
}
