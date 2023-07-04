package models

type TagModel struct {
	Model
	Name          string         `json:"name"`
	Color         string         `json:"color"`
	ArticleModels []ArticleModel `gorm:"many2many:article_tags;joinForeignKey:TagId;joinReferences:ArticleId"`
}
