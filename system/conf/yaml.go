package conf

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var Config Configuration

type Configuration struct{
	Database struct{
		DBName		string `yaml:"dbname"`
		Host		string `yaml:"host"` 
		Username	string `yaml:"username"`
		Password	string `yaml:"password"`
		Type 	  []string `yaml:"type"`
	} `yaml:"database"`
} 

func (c *Configuration) GetConfiguration() {
	yamlFile, err := ioutil.ReadFile("env.yml")
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



