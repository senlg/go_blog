package api

import (
	"go_blog/api"

	"github.com/gin-gonic/gin"
)

func InitCommonApi(apiGroup *gin.RouterGroup) {
	var t api.TokenApi
	// 刷新token
	apiGroup.POST("refresh_toekn", t.RefreshToken)
}
