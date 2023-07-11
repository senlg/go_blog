package api

import (
	"errors"
	"go_blog/common"
	"go_blog/global"
	"go_blog/models"
	"go_blog/models/req"
	"go_blog/models/res"

	"github.com/gin-gonic/gin"
)

type TagApi struct{}

// 查询
func (t *TagApi) GetTagList(ctx *gin.Context) {
	var requestStruct req.TagListRequest
	var TagReponseList []res.TagListItem
	var TagModeles []models.TagModel
	var count int64
	var response common.Response
	ctx.BindJSON(&requestStruct)

	if requestStruct.Limit == 0 {
		requestStruct.Limit = 5
	}
	if requestStruct.Page < 1 {
		requestStruct.Page = 1
	}
	if requestStruct.Order == "" {
		requestStruct.Order = "desc"
	}
	offset := requestStruct.Limit * (requestStruct.Page - 1)

	err := global.DB.Preload("ArticleModels").Model(&models.TagModel{Name: requestStruct.TagName}).Limit(requestStruct.Limit).Offset(offset).Find(&TagModeles).Count(&count).Error

	if err != nil {
		global.Log.Errorln(err.Error())
		response.ResultWithError(ctx, common.RequestError, err)
		return
	}

	for _, v := range TagModeles {
		TagReponseList = append(TagReponseList, res.TagListItem{
			Id:            v.ID,
			TagName:       v.Name,
			Color:         v.Color,
			ArticleCounts: len(v.ArticleModels),
			CreatedAt:     v.CreatedAt,
		})
	}

	response.ResultOk(ctx, common.ListResponse[res.TagListItem]{
		List:  TagReponseList,
		Count: count,
	})
}

// 增加
func (t *TagApi) CreateTag(ctx *gin.Context) {
	// 绑定
	var requestStruct req.TagInfo
	var response common.Response
	ctx.BindJSON(&requestStruct)

	if requestStruct.TagName == "" {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}
	if requestStruct.Color == "" {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}

	tagItem := &models.TagModel{
		Name:  requestStruct.TagName,
		Color: requestStruct.Color,
	}
	global.DB.Create(&tagItem)

	response.ResultOk(ctx, &res.TagListItem{
		Id:        tagItem.ID,
		TagName:   tagItem.Name,
		Color:     tagItem.Color,
		CreatedAt: tagItem.CreatedAt,
	})
}

// 更新
func (t *TagApi) UpdateTag(ctx *gin.Context) {
	// 绑定
	var requestStruct req.TagInfo
	var response common.Response
	ctx.BindJSON(&requestStruct)

	if requestStruct.Id == 0 {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少id"))
		return
	}
	if requestStruct.TagName == "" {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}
	if requestStruct.Color == "" {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}
	var tagItem models.TagModel
	err := global.DB.Find(&tagItem, requestStruct.Id).Error
	if err != nil {
		global.Log.Errorln(err.Error())
		response.ResultWithError(ctx, common.RequestError, err)
		return
	}

	tagItem.Color = requestStruct.Color
	tagItem.Name = requestStruct.TagName
	global.DB.Save(&tagItem)
	response.ResultOk(ctx, "")
}

// 删除
func (t *TagApi) DeleteTag(ctx *gin.Context) {
	// 绑定
	var requestStruct req.DeteleTag
	var response common.Response
	ctx.BindJSON(&requestStruct)

	if requestStruct.Ids == nil || len(requestStruct.Ids) < 1 {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少参数ids"))
		return
	}

	var tagList []models.TagModel
	err := global.DB.Delete(&tagList, requestStruct.Ids).Error
	if err != nil {
		global.Log.Errorln(err.Error())
		response.ResultWithError(ctx, common.RequestError, err)
		return
	}
	response.ResultOk(ctx, "")
}
