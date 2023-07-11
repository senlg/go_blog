package req

type RefreshToken struct {
	OldToken string `json:"old_token"`
}
