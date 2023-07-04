package models

import (
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

type ImageUseType int

type FileUseType int

type ImageModel struct {
	gorm.Model
	Url      string       `gorm:"size:256" json:"url"` //路径
	UseType  ImageUseType `json:"use_type"`
	TypeName string       `json:"type_name"`
	Md5      string       `json:"md5"`
	Enable   string       `json:"enable" gorm:"size:8"` //是否启用
}
type FilesModel struct {
	gorm.Model
	Url      string      `gorm:"size:256" json:"url"` //路径
	UseType  FileUseType `json:"use_type"`
	TypeName string      `json:"type_name"`
	Md5      string      `json:"md5"`
	Enable   string      `json:"enable" gorm:"size:8"` //是否启用
}
