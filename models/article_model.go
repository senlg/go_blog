package models

type ArticleModel struct {
	Model
	UserId     uint           `json:"-"`
	Title      string         `json:"title"`
	Content    string         `json:"content"`
	UserModels []UserModel    `gorm:"many2many:user_collects" json:"-"`
	Tags       []TagModel     `gorm:"many2many:article_tags;" json:"-"`
	Comments   []CommentModel `gorm:"foreignKey:ArticleId"`
}
