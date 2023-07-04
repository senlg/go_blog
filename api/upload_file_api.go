package api

import (
	"errors"
	"fmt"
	"go_blog/common"
	"go_blog/global"
	"go_blog/models"
	modelsRes "go_blog/models/res"
	"go_blog/utils"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UploadFile struct {
	RequestKey string
}

// 上传文件
func (f *UploadFile) Upload(c *gin.Context) {
	res := common.Response{}
	var ResponseList []modelsRes.FileUploadInfo
	if f.RequestKey == "files" {
		ResponseList = UploadFileHandle(f, &res, c)
	} else if f.RequestKey == "images" {
		ResponseList = UploadImageHandle(f, &res, c)
	}

	res.Data = ResponseList
	res.Code = 0
	res.Result(c)
}

// 判断文件大小是否复合限制
func DetermineFileSize(file *multipart.FileHeader) bool {
	return (int64(file.Size) / int64(1024*1024)) < global.Config.UploadConfig.LimitSize
}

func UploadImageHandle(f *UploadFile, res *common.Response, c *gin.Context) (ResponseList []modelsRes.FileUploadInfo) {
	useType := c.PostForm("use_type")
	if useType == "" {
		res.ResultWithError(c, common.RequestError, errors.New("缺少参数"))
		return
	}
	fileList, err := c.MultipartForm()
	if err != nil {
		res.ResultWithError(c, common.UploadError, err)
		return
	}

	var dbfileList []models.ImageModel
	images, ok := fileList.File[f.RequestKey]
	if !ok {
		res.ResultWithError(c, common.UploadError, errors.New("请选择文件上传"))
		return
	}
	_, err = os.ReadDir(global.Config.UploadConfig.BasePath)
	if err != nil {
		err = os.MkdirAll(global.Config.UploadConfig.BasePath, fs.ModePerm)
		if err != nil {
			global.Log.Errorln(err.Error())
		}
	}

	for _, file := range images {
		i := strings.Split(file.Filename, ".")
		fileType := i[len(i)-1]
		e := utils.Find[string](global.Config.UploadConfig.ImgAccessType, fileType)
		if !e {
			ResponseList = append(ResponseList, modelsRes.FileUploadInfo{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("上传文件格式不正确,当前文件格式为%s", fileType),
			})
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			global.Log.Errorln(err.Error())
		}
		byteAr, err := io.ReadAll(fileReader)
		if err != nil {
			global.Log.Errorln(err.Error())
		}
		md5str := utils.Md5(byteAr, "")
		imageM := models.ImageModel{}
		err = global.DB.Take(&imageM, "md5 = ?", md5str).Error
		if err == nil {
			// 找到相同md5图片
			ResponseList = append(ResponseList, modelsRes.FileUploadInfo{
				FileName:  imageM.Url,
				IsSuccess: true,
				Msg:       "图片已存在",
			})
			continue
		}
		if DetermineFileSize(file) {
			filepath := path.Join(global.Config.UploadConfig.BasePath, f.RequestKey, file.Filename)
			err := c.SaveUploadedFile(file, filepath)
			if err != nil {
				ResponseList = append(ResponseList, modelsRes.FileUploadInfo{
					FileName:  filepath,
					IsSuccess: false,
					Msg:       err.Error(),
				})
				continue
			}
			ut, err := strconv.Atoi(useType)
			if err != nil {
				ut = 5
			}
			dbfileList = append(dbfileList, models.ImageModel{
				Url:      filepath,
				UseType:  models.ImageUseType(ut),
				TypeName: fileType,
				Md5:      md5str,
				Enable:   "1",
			})
			ResponseList = append(ResponseList, modelsRes.FileUploadInfo{
				FileName:  filepath,
				IsSuccess: true,
				Msg:       "上传成功",
			})
		} else {
			ResponseList = append(ResponseList, modelsRes.FileUploadInfo{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("文件大小超出限制大小,当前文件大小为%.2dMB,设定大小为%dMB", file.Size, global.Config.UploadConfig.LimitSize),
			})
		}
	}
	if len(dbfileList) > 0 {
		global.DB.Create(&dbfileList)
	}
	return
}
func UploadFileHandle(f *UploadFile, res *common.Response, c *gin.Context) (ResponseList []modelsRes.FileUploadInfo) {
	useType := c.PostForm("use_type")
	if useType == "" {
		res.ResultWithError(c, common.RequestError, errors.New("缺少参数"))
		return
	}
	fileList, err := c.MultipartForm()
	if err != nil {
		res.ResultWithError(c, common.UploadError, err)
		return
	}

	var dbfileList []models.FilesModel
	images, ok := fileList.File[f.RequestKey]
	if !ok {
		res.ResultWithError(c, common.UploadError, errors.New("请选择文件上传"))
		return
	}
	_, err = os.ReadDir(global.Config.UploadConfig.BasePath)
	if err != nil {
		err = os.MkdirAll(global.Config.UploadConfig.BasePath, fs.ModePerm)
		if err != nil {
			global.Log.Errorln(err.Error())
		}
	}

	for _, file := range images {
		i := strings.Split(file.Filename, ".")
		fileType := i[len(i)-1]
		e := utils.Find[string](global.Config.UploadConfig.ImgAccessType, fileType)
		if !e {
			ResponseList = append(ResponseList, modelsRes.FileUploadInfo{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("上传文件格式不正确,当前文件格式为%s", fileType),
			})
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			global.Log.Errorln(err.Error())
		}
		byteAr, err := io.ReadAll(fileReader)
		if err != nil {
			global.Log.Errorln(err.Error())
		}
		md5str := utils.Md5(byteAr, "")
		imageM := models.ImageModel{}
		err = global.DB.Take(&imageM, "md5 = ?", md5str).Error
		if err == nil {
			// 找到相同md5图片
			ResponseList = append(ResponseList, modelsRes.FileUploadInfo{
				FileName:  imageM.Url,
				IsSuccess: true,
				Msg:       "图片已存在",
			})
			continue
		}
		if DetermineFileSize(file) {
			filepath := path.Join(global.Config.UploadConfig.BasePath, f.RequestKey, file.Filename)
			err := c.SaveUploadedFile(file, filepath)
			if err != nil {
				ResponseList = append(ResponseList, modelsRes.FileUploadInfo{
					FileName:  filepath,
					IsSuccess: false,
					Msg:       err.Error(),
				})
				continue
			}
			ut, err := strconv.Atoi(useType)
			if err != nil {
				ut = 5
			}
			dbfileList = append(dbfileList, models.FilesModel{
				Url:      filepath,
				UseType:  models.FileUseType(ut),
				TypeName: fileType,
				Md5:      md5str,
				Enable:   "1",
			})
			ResponseList = append(ResponseList, modelsRes.FileUploadInfo{
				FileName:  filepath,
				IsSuccess: true,
				Msg:       "上传成功",
			})
		} else {
			ResponseList = append(ResponseList, modelsRes.FileUploadInfo{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("文件大小超出限制大小,当前文件大小为%.2dMB,设定大小为%dMB", file.Size, global.Config.UploadConfig.LimitSize),
			})
		}
	}
	if len(dbfileList) > 0 {
		global.DB.Create(&dbfileList)
	}
	return
}
