package api

import (
	"go_blog/api"

	"github.com/gin-gonic/gin"
)

func InitCommoentApi(apiGroup *gin.RouterGroup) {
	var c api.CommontApi
	// 留言
	apiGroup.POST("leave_comment", c.LeaveComment)
	apiGroup.POST("argee_comment", c.AgreeCommont)
}
