package api

import (
	"errors"
	"fmt"
	"go_blog/common"
	"go_blog/global"
	"go_blog/models"
	"go_blog/models/req"
	"go_blog/models/res"
	"go_blog/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ArticleApi struct{}

// 创建文章
func (a *ArticleApi) CreateArticle(ctx *gin.Context) {
	var requestStruct req.ArticleInfo
	var response common.Response
	ctx.BindJSON(&requestStruct)
	if requestStruct.UserId == 0 {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}
	if requestStruct.Title == "" {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}
	if requestStruct.Content == "" {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}
	if requestStruct.TagIds == nil || len(requestStruct.TagIds) < 1 {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}

	var articleModel models.ArticleModel
	var TagModleList []models.TagModel
	err := global.DB.Find(&TagModleList, requestStruct.TagIds).Error
	if err != nil {
		global.Log.Errorln(err.Error())
		response.ResultWithError(ctx, common.RequestError, err)
		return
	}
	articleModel.Content = requestStruct.Content
	articleModel.Title = requestStruct.Title
	articleModel.Tags = TagModleList
	articleModel.UserId = requestStruct.UserId

	err = global.DB.Create(&articleModel).Error
	if err != nil {
		global.Log.Errorln(err.Error())
		response.ResultWithError(ctx, common.RequestError, err)
		return
	}
	response.ResultOk(ctx, "")
}

// 更改文章
func (a *ArticleApi) UpdateArticle(ctx *gin.Context) {
	var requestStruct req.ArticleInfo
	var response common.Response
	var articleModel models.ArticleModel
	var TagModleList []models.TagModel
	ctx.BindJSON(&requestStruct)
	if requestStruct.ArticleId == 0 {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	} else {
		articleModel.ID = requestStruct.ArticleId
	}
	if requestStruct.Title != "" {
		articleModel.Title = requestStruct.Title
	}
	if requestStruct.Content == "" {
		articleModel.Content = requestStruct.Content
	}
	if requestStruct.TagIds != nil && len(requestStruct.TagIds) > 0 {
		err := global.DB.Find(&TagModleList, requestStruct.TagIds).Error
		if err != nil {
			global.Log.Errorln(err.Error())
			response.ResultWithError(ctx, common.RequestError, err)
			return
		}
		err = global.DB.Model(&articleModel).Association("Tags").Replace(TagModleList)
		if err != nil {
			global.Log.Errorln(err.Error())
			response.ResultWithError(ctx, common.RequestError, err)
			return
		}
	}

	err := global.DB.Preload("Tags").Where("id =?", requestStruct.ArticleId).Find(&articleModel).Error
	global.DB.Save(&articleModel)
	fmt.Printf("articleModel: %+v\n", articleModel)
	if err != nil {
		global.Log.Errorln(err.Error())
		response.ResultWithError(ctx, common.RequestError, err)
		return
	}
	response.ResultOk(ctx, "")
}

// 删除文章
func (a *ArticleApi) DeleteArticle(ctx *gin.Context) {
	var requestStruct req.ArticleInfo
	var response common.Response
	var articleModel models.ArticleModel
	ctx.BindJSON(&requestStruct)
	if requestStruct.ArticleId == 0 {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	} else {
		articleModel.ID = requestStruct.ArticleId
	}
	err := global.DB.Select(clause.Associations).Delete(&articleModel).Error
	if err != nil {
		response.ResultWithError(ctx, common.ErrorStatus, err)
		return
	}
	response.ResultOk(ctx, "")
}

// 查询文章列表
func (a *ArticleApi) GetArticleList(ctx *gin.Context) {
	var requestStruct req.ArticleListRequest
	var response common.Response
	var articleModels []models.ArticleModel
	var tx *gorm.DB = global.DB.Model(&models.ArticleModel{})
	var count int64
	ctx.BindJSON(&requestStruct)
	if requestStruct.Limit < 5 {
		requestStruct.Limit = 5
	}
	if requestStruct.Page < 1 {
		requestStruct.Page = 1
	}
	if requestStruct.CreateDateEnd != "" && requestStruct.CreateDateStart != "" {
		ok, err := utils.VerificationDateList([]string{requestStruct.CreateDateStart, requestStruct.CreateDateEnd})
		if ok {
			tx = tx.Where("created_at > ? and created_at < ?", requestStruct.CreateDateStart, requestStruct.CreateDateEnd)
		} else {
			response.ResultWithError(ctx, common.RequestError, err)
		}
	}
	if requestStruct.Order != "" {
		s := fmt.Sprintf("created_at %s", requestStruct.Order)
		tx = tx.Order(s)
	} else {
		tx = tx.Order("created_at desc")
	}

	if requestStruct.Title != "" {
		tx = tx.Where(fmt.Sprintf("title like '%%%s%%' ", requestStruct.Title))
	}
	tx = tx.Preload("Tags")
	if len(requestStruct.TagIds) > 0 {
		tx = tx.Joins("JOIN article_tags at ON at.article_id = article_models.id JOIN tag_models tm ON tm.id = at.tag_id").Where("tm.id in ?", requestStruct.TagIds).Group("at.article_id")
		err := tx.Error
		if err != nil {
			response.ResultWithError(ctx, common.ErrorStatus, err)
			return
		}
	}
	tx = tx.Preload("UserModels")
	offset := requestStruct.Limit * (requestStruct.Page - 1)
	tx.Limit(requestStruct.Limit).Offset(offset).Find(&articleModels).Count(&count)
	var articleList []res.ArticleItem
	for _, v := range articleModels {
		var tags []res.Tag
		var user models.UserModel
		for _, t := range v.Tags {
			tags = append(tags, res.Tag{
				Id:      t.ID,
				TagName: t.Name,
				Color:   t.Color,
			})
		}
		global.DB.Take(&user, v.UserId)
		articleList = append(articleList, res.ArticleItem{
			ID:        v.ID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Title:     v.Title,
			Tags:      tags,
			UserId:    v.UserId,
			UserName:  user.UserName,
		})

	}
	response.ResultOk(ctx, common.ListResponse[res.ArticleItem]{
		List:  articleList,
		Count: count,
	})
}
