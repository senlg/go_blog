package models

import (
	"go_blog/global"
	"os"

	"gorm.io/gorm"
)

// 图片类型
const (
	ImageAvatar  ImageUseType = iota + 1 //头像
	ImageBanner                          //轮播图
	ImageArticle                         //文章图片
	ImageCurrent                         //通用
	ImageError                           // 上传解析错误时
)

// 文件类型
const (
	FileMd    FileUseType = iota + 1 // markdown
	FileOther                        // 其他
)

type FileSource int

const (
	FileTypeLocal FileSource = iota + 1 // 本地
	FileTypeQiniu FileSource = iota + 1 // 七牛
)

// 图片用途
type ImageUseType int

// 文件用途
type FileUseType int

type ImageModel struct {
	gorm.Model
	Url        string       `gorm:"size:256" json:"url"`          // 路径
	UseType    ImageUseType `json:"use_type"`                     // 用途
	TypeName   string       `json:"type_name"`                    // 文件后缀名
	FileName   string       `json:"file_name"`                    // 文件名
	Md5        string       `json:"md5"`                          // 内容md5
	Enable     string       `json:"enable" gorm:"size:8"`         // 是否启用
	FileSource FileSource   `json:"file_source" gorm:"default:1"` // 本地还是七牛
}
type FilesModel struct {
	gorm.Model
	Url        string      `gorm:"size:256" json:"url"` //路径
	UseType    FileUseType `json:"use_type"`
	TypeName   string      `json:"type_name"`
	FileName   string      `json:"file_name"` // 文件名
	Md5        string      `json:"md5"`
	Enable     string      `json:"enable" gorm:"size:8"`         //是否启用
	FileSource FileSource  `json:"file_source" gorm:"default:1"` // 本地还是七牛
}

func (i *ImageModel) AfterDelete(tx *gorm.DB) (err error) {
	if i.FileSource == FileTypeLocal {
		err = os.Remove(i.Url)
		if err != nil {
			global.Log.Errorln(err.Error())
		}
	}
	return err
}
