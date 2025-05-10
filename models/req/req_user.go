package req

type RegisterUser struct {
	UserName  string `json:"user_name"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	Addr      string `json:"addr"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type Login struct {
	UserName string `json:"user_name"` // 用户名
	Password string `json:"password"`  // 密码
	// Captcha     string `json:"captcha"`      // 验证码
	LoginAdress string `json:"login_adress"` //登录地址
}

type UserInfo struct {
	UserId uint `json:"user_id"`
}
type UserInfoRequest struct {
	// UserId   uint   `json:"user_id"`
	UserName string `json:"user_name"` // 用户名
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
	// NickName string `json:"nick_name"` 	// 昵称
}
type DeleteUserInfo struct {
	Ids []uint `json:"ids"`
}

// 修改密码
type ChangePassWord struct {
	UserName    string `json:"user_name"` // 用户名
	OldPassword string `json:"old_password"`
	Password    string `json:"password"` // 密码
	Code        string `json:"code"`
}
