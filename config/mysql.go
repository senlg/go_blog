package config

import "fmt"

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Config   string `yaml:"config"`  //高级配置 例如charset
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` // 日志等级
}

func (m *Mysql) GetDsn() (dsn string) {
	dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", m.User,m.Password,m.Host,m.Port,m.Db,m.Config)
	return
}