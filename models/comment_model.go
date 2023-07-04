package models

type CommentModel struct {
	Model
	Name            string         `json:"name"`
	Comment         string         `json:"comment"`
	MainId          *uint          `json:"main_id"`
	ChildrenComment []CommentModel `gorm:"foreignKey:MainId"`
	ArticleId       uint           `json:"article_id"`
}
