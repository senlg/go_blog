package imageserver

import (
	"fmt"
	"go_blog/global"
	"go_blog/models"
	"go_blog/models/res"
	"go_blog/plugins/qiniu"
	"go_blog/utils"
	"io"
	"mime/multipart"
	"strconv"
	"strings"
)

type ImageServer struct {
}

func (ImageServer) UploadSingleFileServer(file *multipart.FileHeader, params map[string]string) (FileUploadInfo res.FileUploadInfo, dbFileInfo models.ImageModel) {
	useType := params["useType"]
	fileSource := params["fileSource"]

	// responseInfo可初始化
	FileUploadInfo.FileName = file.Filename
	FileUploadInfo.Msg = "未知错误"

	// DBFileInfo 可初始化
	ut, err := strconv.Atoi(useType)
	if err != nil {
		ut = 5
	}
	// 截取文件后缀名
	i := strings.Split(file.Filename, ".")
	fileType := i[len(i)-1]
	// 计算md5
	fileReader, err := file.Open()
	if err != nil {
		global.Log.Errorln(err.Error())
	}
	byteAr, err := io.ReadAll(fileReader)
	if err != nil {
		global.Log.Errorln(err.Error())
	}
	md5str := utils.Md5(byteAr, "")
	dbFileInfo.UseType = models.ImageUseType(ut)
	dbFileInfo.TypeName = fileType
	dbFileInfo.Md5 = md5str
	dbFileInfo.Enable = "1"

	e := utils.Find[string](global.Config.UploadConfig.ImgAccessType, fileType)
	if !e {
		FileUploadInfo.IsSuccess = false
		FileUploadInfo.Msg = fmt.Sprintf("上传文件格式不正确,当前文件格式为%s", fileType)
		return
	}

	imageM := models.ImageModel{}
	err = global.DB.Take(&imageM, "md5 = ?", md5str).Error
	if err == nil {
		// 找到相同md5图片
		FileUploadInfo.FileUrl = imageM.Url
		FileUploadInfo.IsSuccess = false
		FileUploadInfo.Msg = "图片已存在"
		return
	}
	if fileSource == "" {
		fileSource = string(rune(models.FileTypeLocal))
	}
	if fileSource == "2" {
		filePath, err := qiniu.UploadQiNiu(byteAr, file.Filename, "go_blog")
		if err != nil {
			global.Log.Error(err.Error())
			return
		}
		dbFileInfo.Url = filePath
		FileUploadInfo.FileUrl = filePath
		FileUploadInfo.IsSuccess = true
		FileUploadInfo.Msg = "上传成功"
		return
	}

	return
}
