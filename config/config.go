package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

//Config 配置结构
type Config struct {
	App struct {
		Host  string `yaml:"host"`
		Debug bool   `yaml:"debug"`
	}
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
			Debug  bool   `yaml:"debug"`
		}
		Redis struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		}
	}
	Wechat struct {
		User struct {
			AppID     string `yaml:"appid"`
			AppSecret string `yaml:"app-secret"`
		}
	}
}

//Info 全局配置
var Info Config

func init() {
	getInfo(&Info, "/config/config.yaml")
}

func getInfo(conf *Config, file string) {
	dir, _ := os.Getwd()
	configFile, err := ioutil.ReadFile(dir + file)
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	err = yaml.Unmarshal(configFile, &conf)
	if err != nil {
		log.Panicln("Unmarshal:", err.Error())
	}
}
