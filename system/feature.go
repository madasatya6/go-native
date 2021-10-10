package system

import (
	"github.com/gorilla/mux"
	"github.com/madasatya6/go-native/system/route"
	"github.com/madasatya6/go-native/system/env"
	"github.com/madasatya6/go-native/system/logs"
)

//Register your feature here
type Features struct {
	Route 		*mux.Router
	Env			*env.Environment
	LogEntry	*logs.Log
} 

func Init() *Features {
	var ft Features
	ft.Route	 = route.Init()
	ft.Env		 = env.Init()
	ft.LogEntry	 = logs.Init()
	return &ft
}

