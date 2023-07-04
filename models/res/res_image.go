package res

import "go_blog/models"

type ImageInfo struct {
	Url        string              `json:"url"`
	UseType    models.ImageUseType `json:"use_type"`
	ImageName  string              `json:"image_name"`
	CreateDate string              `json:"create_date"`
	TypeName   string              `json:"type_name"`
	Enable     string              `json:"enable"`
	Id         uint                `json:"id"`
}
