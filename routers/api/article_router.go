package api

import (
	"go_blog/api"

	"github.com/gin-gonic/gin"
)

func InitArticleApi(apiGroup *gin.RouterGroup) {
	var a api.ArticleApi
	var c api.CollectApi
	apiGroup.POST("create_article", a.CreateArticle)
	apiGroup.POST("update_article", a.UpdateArticle)
	apiGroup.POST("get_articles", a.GetArticleList)
	apiGroup.DELETE("delete_article", a.DeleteArticle)
	apiGroup.POST("cancel_article", c.DelArticle)
	apiGroup.POST("collect_article", c.CollectArticle)
}
