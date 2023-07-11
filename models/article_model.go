package models

type ArticleModel struct {
	Model
	UserId     uint           `json:"user_id" gorm:"column:user_id"`
	Title      string         `json:"title" gorm:"size: 128"`
	Content    string         `json:"content"`
	UserModels []UserModel    `gorm:"many2many:user_collects" json:"collect_users"`
	Tags       []TagModel     `gorm:"many2many:article_tags;joinForeignKey:ArticleId;joinReferences:TagId" json:"tags"`
	Comments   []CommentModel `gorm:"foreignKey:ArticleId" json:"commonts"`
}
