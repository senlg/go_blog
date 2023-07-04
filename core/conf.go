package core

import (
	"fmt"
	"go_blog/config"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func InitConf() *config.Config {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}

	yamlConf, err := ioutil.ReadFile(ConfigFile)

	if err != nil {
		panic(fmt.Errorf("get yamlFile error: %v", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("Config Init Unmarshal Error: %v", err)
	}
	log.Printf("Config yaml File load init success!")
	return c
}
