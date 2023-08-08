package api

import (
	"errors"
	"go_blog/common"
	"go_blog/global"
	"go_blog/models"
	"go_blog/models/req"

	"github.com/gin-gonic/gin"
)

type CommontApi struct {
}

// 留言or回复
func (c CommontApi) LeaveComment(ctx *gin.Context) {
	var requestStruct req.LeaveCommentRequest
	ctx.BindJSON(&requestStruct)
	var response common.Response
	// Comment

	// ArticleId
	// UserId
	// ReplyCommentId
	// ReplyUserId
	// 回复

	if requestStruct.ArticleId == 0 {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}
	if requestStruct.Comment == "" {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}
	if requestStruct.UserId == 0 {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, requestStruct.UserId).Error
	if err != nil {
		response.ResultWithError(ctx, common.RequestError, errors.New("未找到user_id相关用户"))
		return
	}
	// 留言
	if requestStruct.ReplyMainCommentId == 0 && requestStruct.ReplyUserId == 0 {
		global.DB.Create(&models.CommentModel{
			Comment:   requestStruct.Comment,
			ArticleId: requestStruct.ArticleId,
			UserModel: user,
		})
		response.ResultOk(ctx, "")
		return

	} else { // 回复
		var main_model models.CommentModel
		main_model.ID = requestStruct.ReplyMainCommentId

		err := global.DB.Model(&main_model).Association("ChildrenComment").Append(&models.CommentModel{
			Comment:     requestStruct.Comment,
			MainId:      &requestStruct.ReplyMainCommentId,
			ArticleId:   requestStruct.ArticleId,
			UserId:      requestStruct.UserId,
			ReplyUserId: requestStruct.ReplyUserId,
		})
		if err != nil {
			response.ResultWithError(ctx, common.ErrorStatus, err)
			return
		}
		response.ResultOk(ctx, "")
		return
	}
}

// 点赞

func (c CommontApi) AgreeCommont(ctx *gin.Context) {
	var requestStruct req.CommentAgree
	ctx.BindJSON(&requestStruct)
	var response common.Response

	if requestStruct.UserId == 0 {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}
	if requestStruct.CommentId == 0 {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}
	if !requestStruct.IsAgree {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}

	var commont_model = models.CommentModel{}
	commont_model.ID = requestStruct.CommentId
	err := global.DB.Find(&commont_model).Association("AgreeModels").Append(&models.AgreeModel{
		IsAgree:   requestStruct.IsAgree,
		CommentId: requestStruct.CommentId,
		UserId:    requestStruct.UserId,
	})
	if err != nil {
		response.ResultWithError(ctx, common.ErrorStatus, err)
		return
	}
	response.ResultOk(ctx, "")

}
