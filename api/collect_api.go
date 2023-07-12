package api

import (
	"errors"
	"fmt"
	"go_blog/common"
	"go_blog/global"
	"go_blog/models"
	"go_blog/models/req"

	"github.com/gin-gonic/gin"
)

type CollectApi struct {
}

func (c *CollectApi) CollectArticle(ctx *gin.Context) {
	var response common.Response
	var resquestStruct req.CollectArticleRequset
	ctx.BindJSON(&resquestStruct)

	if resquestStruct.CollectId == 0 {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少文章id"))
		return
	}
	if resquestStruct.UserId == 0 {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少用户id"))
		return
	}
	article := models.ArticleModel{}
	affect := global.DB.Find(&article, resquestStruct.CollectId).RowsAffected
	if affect < 1 {
		response.ResultWithError(ctx, common.RequestError, errors.New("无相关文章可收藏"))
		return
	} else if affect > 1 {
		response.ResultWithError(ctx, common.ErrorStatus, errors.New("数据库数据错误"))
		return
	}

	var user models.UserModel
	global.DB.Model(&article).Association("UserModels").Find(&user, resquestStruct.UserId)

	if user.ID == 0 {
		err := global.DB.Find(&user, resquestStruct.UserId).Error
		if err != nil {
			response.ResultWithError(ctx, common.ErrorStatus, errors.New("数据库数据错误"))
			return
		}
		fmt.Printf("userIds: %v\n", user)
		global.DB.Model(&article).Association("UserModels").Append(&user)
		response.Msg = "收藏成功"
		response.ResultOk(ctx, "")
		return
	} else {
		response.ResultWithError(ctx, common.RequestError, errors.New("已收藏文章"))
	}

}
func (c *CollectApi) DelArticle(ctx *gin.Context) {
	var response common.Response
	var resquestStruct req.CollectArticleRequset
	ctx.BindJSON(&resquestStruct)

	if resquestStruct.CollectId == 0 {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少收藏文章id"))
		return
	}
	if resquestStruct.UserId == 0 {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少用户id"))
		return
	}
	article := models.ArticleModel{}
	affect := global.DB.Find(&article, resquestStruct.CollectId).RowsAffected
	fmt.Printf("affect: %v\n", affect)
	if affect < 1 {
		response.ResultWithError(ctx, common.RequestError, errors.New("无相关文章"))
		return
	} else if affect > 1 {
		response.ResultWithError(ctx, common.ErrorStatus, errors.New("数据库数据错误"))
		return
	}
	fmt.Printf("article: %v\n", article)
	var user models.UserModel
	user.ID = resquestStruct.UserId
	err := global.DB.Model(&article).Association("UserModels").Delete(user)
	fmt.Printf("users: %v\n", user)
	fmt.Printf("err: %+v\n", err)
	response.Msg = "取消成功"
	response.ResultOk(ctx, "")
}
