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
	log.Printf("\n Config yaml File load init success! \n")
	return c
}

// func InitIntercept() *config.InterceptApiYaml {
// 	const InterceptApiFile = "routers/intercept_api.yaml"
// 	i := &config.InterceptApiYaml{}
// 	yamlConf, err := ioutil.ReadFile(InterceptApiFile)
// 	if err != nil {
// 		panic(fmt.Errorf("get yamlFile error: %v", err))
// 	}
// 	err = yaml.Unmarshal(yamlConf, i)
// 	if err != nil {
// 		log.Fatalf("Config Init Unmarshal Error: %v", err)
// 	}
// 	log.Printf("InterceptApi yaml File load init success! \n")
// 	return i
// }
