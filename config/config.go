package config

type Config struct {
	Mysql            Mysql            `yaml:"mysql"`
	Logger           Logger           `yaml:"logger"`
	System           System           `yaml:"system"`
	Qiniu            Qiniu            `yaml:"qiniu"`
	UploadConfig     UploadConfig     `yaml:"upload_config"`
	InterceptApiYaml InterceptApiYaml `yaml:"intercept_router"`
}

// 拦截校验jwt的路由
type InterceptApiYaml struct {
	InterceptPath []string `yaml:"intercept_path"`
}
