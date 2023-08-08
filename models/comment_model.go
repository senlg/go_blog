package models

type CommentModel struct {
	Model
	Comment         string         `json:"comment"`
	MainId          *uint          `json:"main_id"`
	ChildrenComment []CommentModel `gorm:"foreignKey:MainId"`
	ArticleId       uint           `json:"article_id"`
	UserId          uint           `json:"user_id"`
	UserModel       UserModel      `gorm:"foreignKey:UserId"`
	ReplyUserId     uint           `json:"reply_user_id"`
	AgreeCount      uint           `json:"agree_count"`
	AgreeModels     []AgreeModel   `gorm:"foreignKey:CommentId"`
}

type AgreeModel struct {
	ID        uint `json:"id"`
	IsAgree   bool `json:"is_agree"`
	CommentId uint `json:"comment_id"`
	UserId    uint `json:"user_id"`
}
