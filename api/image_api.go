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
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ImageApi struct{}

// 获取列表
func (i *ImageApi) FindImageList(ctx *gin.Context) {
	var ImageSearchStruct req.GetImageReq

	ctx.BindJSON(&ImageSearchStruct)
	fmt.Printf("%+v\n", ImageSearchStruct)

	var imagesModelList []models.ImageModel
	var imagesList []res.ImageInfo
	var count int64
	var tx *gorm.DB = global.DB.Model(&models.ImageModel{})
	response := common.Response{}
	if ImageSearchStruct.Limit < 5 {
		ImageSearchStruct.Limit = 5
	}
	if ImageSearchStruct.Page < 1 {
		ImageSearchStruct.Page = 1
	}
	if ImageSearchStruct.CreateDateEnd != "" && ImageSearchStruct.CreateDateStart != "" {
		ok, err := utils.VerificationDateList([]string{ImageSearchStruct.CreateDateStart, ImageSearchStruct.CreateDateEnd})
		if ok {
			tx = tx.Where("created_at > ? and created_at < ?", ImageSearchStruct.CreateDateStart, ImageSearchStruct.CreateDateEnd)
		} else {
			response.ResultWithError(ctx, common.RequestError, err)
		}
	}
	if ImageSearchStruct.ImageUrl != "" {
		tx = tx.Where("url =?", ImageSearchStruct.ImageUrl)
	}
	if ImageSearchStruct.TypeName != "" {
		tx = tx.Where("type_name =?", ImageSearchStruct.TypeName)
	}
	if ImageSearchStruct.Enable != "" {
		s := strings.Split(ImageSearchStruct.Enable, ",")
		for _, v := range s {
			fmt.Println(v)
			tx = tx.Where("enable =?", v)
		}
	} else {
		tx = tx.Where("enable = ? or enable = ?", "1", "0")
	}
	if ImageSearchStruct.Order != "" {
		s := fmt.Sprintf("created_at %s", ImageSearchStruct.Order)
		tx = tx.Order(s)
	} else {
		tx = tx.Order("created_at desc")
	}
	tx.Limit(ImageSearchStruct.Limit).Offset((ImageSearchStruct.Page - 1) * ImageSearchStruct.Limit).Find(&imagesModelList).Count(&count)

	for _, v := range imagesModelList {

		imagesList = append(imagesList, res.ImageInfo{
			Id:         v.ID,
			Url:        v.Url,
			UseType:    v.UseType,
			ImageName:  v.FileName,
			CreateDate: v.CreatedAt.Format("2006-01-02 15:04:05"),
			TypeName:   v.TypeName,
			Enable:     v.Enable,
		})
	}
	response.ResultOk(ctx, common.ListResponse[res.ImageInfo]{
		List:  imagesList,
		Count: count,
	})
}

// 更改图片信息
func (i *ImageApi) ChangeImageInfo(ctx *gin.Context) {
	var requestStrut req.ChangeImg
	var imageModel models.ImageModel
	var response common.Response
	ctx.BindJSON(&requestStrut)

	err := global.DB.Take(&imageModel, "id = ?", requestStrut.Id).Error
	if err != nil {

		response.Msg = err.Error()
		response.ResultOk(ctx, "")
		return
	}
	imageModel.Enable = requestStrut.Enable
	if requestStrut.FileName != "" {
		imageModel.FileName = strings.Join([]string{requestStrut.FileName, imageModel.TypeName}, ".")
	}
	err = global.DB.Save(&imageModel).Error
	if err != nil {
		response.Msg = err.Error()
		response.ResultOk(ctx, "")
		return
	}

	response.ResultOk(ctx, "")
}

// 删除图片
func (i *ImageApi) DeleteImage(ctx *gin.Context) {

	var idsStruct req.DeleteImage
	ctx.BindJSON(&idsStruct)
	var imgModel []models.ImageModel
	var count int64
	global.DB.Find(&imgModel, idsStruct.Ids).Count(&count)

	response := common.Response{}
	if count == 0 {
		if global.DB.Error != nil {
			response.ResultWithError(ctx, common.RequestError, global.DB.Error)
			return
		}
		response.ResultWithError(ctx, common.RequestError, errors.New("图片不存在"))
		return
	}
	global.DB.Delete(&imgModel)
	response.Msg = fmt.Sprintf("共删除%d张图片", count)
	response.ResultOk(ctx, "")
}
