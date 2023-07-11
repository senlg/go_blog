package req

type TagListRequest struct {
	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
	Order   string `json:"order"` // desc asc 降序 升序 根据日期
	TagName string `json:"tag_name"`
}

type TagInfo struct {
	Id      uint   `json:"id"`
	TagName string `json:"tag_name"`
	Color   string `json:"color"`
}

type DeteleTag struct {
	Ids []uint `json:"ids"`
}
