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

	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
