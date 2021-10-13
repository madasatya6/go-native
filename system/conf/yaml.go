package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var Config Configuration

type Configuration struct{
	Database struct{
		MySQL struct{
			DBName		string `yaml:"dbname"`
			Host		string `yaml:"host"` 
			Username	string `yaml:"username"`
			Password	string `yaml:"password"`
			TCP 		string `yaml:"tcp"`
		} `yaml:"mysql"`
		Postgre struct{
			DBName		string `yaml:"dbname"`
			Host		string `yaml:"host"` 
			Username	string `yaml:"username"`
			Password	string `yaml:"password"`
			SSLMode 	string `yaml:"sslmode"`
			TCP 		string `yaml:"tcp"`
		} `yaml:"postgre"`
		Type 	  []string `yaml:"type"`
	} `yaml:"database"`
	ToDSN map[string]interface{}
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

func (c *Configuration) SetDSN() {

	MySQL := c.Database.MySQL 
	Postgre := c.Database.Postgre

	mysql := fmt.Sprintf("%v:%v@%s/%v", MySQL.Username, MySQL.Password, MySQL.TCP, MySQL.DBName)
	postgre := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=%s", Postgre.DBName, Postgre.Username, Postgre.Password, Postgre.Host, Postgre.SSLMode)
	c.ToDSN = map[string]interface{}{
		"mysql" : mysql,
		"postgre" : postgre,
	}
}

func Init() *Configuration {
	Config.GetConfiguration()
	Config.SetDSN()
	return &Config
}



