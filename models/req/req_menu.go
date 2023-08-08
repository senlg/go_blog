package req

import (
	"go_blog/models"
)

type CreateMenuRequest struct {
	UserId   uint            `json:"user_id"`
	UserName string          `json:"user_name"`
	MenuName string          `json:"menu_name"`
	MenuType models.MenuType `json:"menu_type"`
	ParentId *uint           `json:"parent_id"`
	// Children []MenuItem      `json:"children"`
}

type MenuItem struct {
	MenuName string          `json:"menu_name"`
	MenuType models.MenuType `json:"menu_type"`
	ParentId *uint           `json:"parent_id"`
	Children []MenuItem      `json:"children"`
}
