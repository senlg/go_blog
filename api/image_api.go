package api

import (
	"fmt"
	"go_blog/common"
	"go_blog/global"
	"go_blog/models"
	"go_blog/models/req"
	"go_blog/models/res"
	"go_blog/utils"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ImageApi struct{}

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
		tx = tx.Where("enable = ? and enable = ?", "1", "0")
	}
	tx.Limit(ImageSearchStruct.Limit).Offset((ImageSearchStruct.Page - 1) * ImageSearchStruct.Limit).Find(&imagesModelList).Count(&count)

	for _, v := range imagesModelList {
		_, ImageName := path.Split(v.Url)
		imagesList = append(imagesList, res.ImageInfo{
			Id:         v.ID,
			Url:        v.Url,
			UseType:    v.UseType,
			ImageName:  ImageName,
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

func (i *ImageApi) ChangeImageEnable(ctx *gin.Context) {
	// ctx.
}
