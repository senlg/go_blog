package req

type ArticleInfo struct {
	UserId    uint   `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	TagIds    []uint `json:"tag_ids"`
	ArticleId uint   `json:"id"`
}

type ArticleListRequest struct {
	Limit           int    `json:"limit"`
	Page            int    `json:"page"`
	Order           string `json:"order"` // desc asc 降序 升序 根据日期
	CreateDateStart string `json:"create_date_start"`
	CreateDateEnd   string `json:"create_date_end"`
	Title           string `json:"title"`
	TagIds          []int  `json:"tag_ids"`
}

type CollectArticleRequset struct {
	CollectId uint `json:"collect_id"`
	UserId    uint `json:"user_id"`
}
