package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Config Configuration

type Configuration struct{
	DBName		string `yaml:"dbname"`
	Host		string `yaml:"host"` 
	Username	string `yaml:"username"`
	Password	string `yaml:"password"`
} 

func (c *Configuration) GetConfiguration() {
	yamlFile, err := ioutil.ReadFile("env.yaml")
	if err != nil {
		log.Println("Yaml file error: ", err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Init() *Configuration {
	Config.GetConfiguration()
	return &Config
}



