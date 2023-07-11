package res

import "go_blog/models"

type TagListItem struct {
	Id            uint         `json:"id"`
	TagName       string       `json:"tag_name"`
	Color         string       `json:"color"`
	ArticleCounts int          `json:"articles"`
	CreatedAt     models.XTime `json:"created_at"`
}
