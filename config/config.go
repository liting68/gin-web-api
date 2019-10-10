package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

//Config 配置结构
type Config struct {
	Middleware struct {
		Secret string `yaml:"secret"`
	}
	DB struct {
		Mysql struct {
			User   string `yaml:"user"`
			Pass   string `yaml:"pass"`
			Host   string `yaml:"host"`
			Port   string `yaml:"port"`
			DBName string `yaml:"dbname"`
		}
	}
}

//Info 全局配置
var Info Config

func init() {
	dir, _ := os.Getwd()
	configFile, err := ioutil.ReadFile(dir + "/config/config.yaml")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	err = yaml.Unmarshal(configFile, &Info)
	if err != nil {
		log.Panicln("Unmarshal:", err.Error())
	}
}
