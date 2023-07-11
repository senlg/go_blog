package res

import (
	"go_blog/models"
)

type ArticleItem struct {
	ID        uint         `json:"id"`
	CreatedAt models.XTime `json:"created_at"`
	UpdatedAt models.XTime `json:"updated_at"`
	Title     string       `json:"title"`
	Tags      []Tag        `json:"tags"`
	UserId    uint         `json:"user_id"`
	UserName  string       `json:"user_name"`
}

type Tag struct {
	Id      uint   `json:"id"`
	TagName string `json:"tag_name"`
	Color   string `json:"color"`
}
