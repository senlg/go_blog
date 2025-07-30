package config

import "fmt"

type System struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Env    string `yaml:"env"`
	Secret string `yaml:"secret"`
}

// 获取监听地址
func (s *System) GetAdress() string {
	fmt.Printf("server listen in http://%s:%d",s.Host, s.Port)
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
