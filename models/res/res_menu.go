package res

import "go_blog/models"

type MenuItem struct {
	MenuId   uint            `json:"menu_id"`
	MenuName string          `json:"menu_name"`
	MenuType models.MenuType `json:"menu_type"`
	ParentId *uint           `json:"parent_id"`
	Children []MenuItem      `json:"children"`
}
