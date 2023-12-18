package config

type Config struct {
	Mysql        Mysql        `yaml:"mysql"`
	Logger       Logger       `yaml:"logger"`
	System       System       `yaml:"system"`
	Qiniu        Qiniu        `yaml:"qiniu"`
	UploadConfig UploadConfig `yaml:"upload_config"`
	InterceptApi InterceptApi `yaml:"intercept_router"`
}

// 拦截校验jwt的路由
type InterceptApi struct {
	InterceptPath []string `yaml:"intercept_path"`
}
