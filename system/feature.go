package system

import (
	"github.com/gorilla/mux"
	"github.com/madasatya6/go-native/system/route"
	"github.com/madasatya6/go-native/system/env"
	"github.com/madasatya6/go-native/system/logs"
	"github.com/madasatya6/go-native/system/database"
	"github.com/madasatya6/go-native/system/conf"
	"github.com/madasatya6/go-native/system/times"
	"github.com/madasatya6/go-native/system/session"
)

//Register your feature here
type Features struct {
	Route 		*mux.Router
	Env			*env.Environment
	LogEntry	*logs.Log
	Databases	*database.Databases
	Conf 		*conf.Configuration
	Time		*times.TimeConf
	Session 	session.SessionType
} 

func Init() *Features {
	var ft Features
	ft.Route	 = route.Init()
	ft.Env		 = env.Init()
	ft.LogEntry	 = logs.Init()
	ft.Conf		 = conf.Init()
	ft.Databases = database.Init(ft.Conf.ToDSN)
	ft.Time	 	 = times.Init()
	ft.Session	 = session.Init()
	return &ft
}

