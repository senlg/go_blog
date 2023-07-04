package config

type Qiniu struct {
	AccessKey string `yaml:"access_key" json:"access_key"`
	SecretKey string `yaml:"secret_key" json:"secret_key"`
	Bucket    string `yaml:"bucket" json:"bucket"` // 存储空间名称 senlg
	Cdn       string `yaml:"cdn" json:"cdn"`       // 访问图片地址前缀
	Zone      string `yaml:"zone" json:"zone"`     // 存储地区
	Size      int    `yaml:"size" json:"size"`     // 存储大小限制 单位mb
}
