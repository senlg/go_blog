package req

type RegisterUser struct {
	UserName  string `json:"name"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	Addr      string `json:"addr"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type Login struct {
	UserName    string `json:"user_name"`    // 用户名
	Password    string `json:"password"`     // 密码
	Captcha     string `json:"captcha"`      // 验证码
	LoginAdress string `json:"login_adress"` //登录地址
}