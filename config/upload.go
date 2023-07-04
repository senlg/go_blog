package config

type UploadConfig struct {
	BasePath      string   `yaml:"base_path"`
	LimitSize     int64    `yaml:"limit_size"`
	ImgAccessType []string `yaml:"img_access_type"`
}
