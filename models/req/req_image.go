package req

import "go_blog/models"

type GetImageReq struct {
	Limit           int    `json:"limit"`
	Page            int    `json:"page"`
	ImageUrl        string `json:"image_url"`
	CreateDateStart string `json:"create_date_start"`
	CreateDateEnd   string `json:"create_date_end"`
	TypeName        string `json:"type_name"`
	Enable          string `json:"enable"` //是否启用 0不启用 1启用
}

type ChangeImg struct {
	Enable  string              `json:"enable"`
	Id      uint                `json:"id"`
	UseType models.ImageUseType `json:"use_type"`
}
