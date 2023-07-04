package models

type MenuType int

type MenuModel struct {
	Model
	MenuName     string `gorm:"size:64" json:"menu_name"`
	MenuType     `json:"menu_type"`
	ParentId     *uint       `json:"parent_id"`
	ChildrenMenu []MenuModel `gorm:"foreignKey:ParentId"`
}
