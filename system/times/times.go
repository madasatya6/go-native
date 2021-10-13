package times

import (
	"time"
	"log"
	"github.com/madasatya6/go-native/applications/config"
)

var Time TimeConf

type TimeConf struct{
	Location string
	TimeZone *time.Location
}

func (t *TimeConf) SetTime(location string){
	t.Location = location
	timezone, err := time.LoadLocation(location)
	if err != nil {
		log.Println(err.Error())
	} 
	t.TimeZone = timezone 
}

func Init() *TimeConf {
	Time.SetTime(config.DefaultLocation)
	config.TimeZone = Time.TimeZone
	return &Time 
}


