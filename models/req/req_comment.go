package req

type LeaveCommentRequest struct {
	Comment            string `json:"comment"`
	ReplyMainCommentId uint   `json:"reply_main_comment_id"`
	ArticleId          uint   `json:"article_id"`
	UserId             uint   `json:"user_id"`
	ReplyUserId        uint   `json:"reply_user_id"`
}

type CommentAgree struct {
	CommentId uint `json:"comment_id"`
	UserId    uint `json:"user_id"`
	IsAgree   bool `json:"is_agree"`
}
