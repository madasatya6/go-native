package system

import (
	"net/http"
	"github.com/madasatya6/go-native/system/route"
	"github.com/madasatya6/go-native/system/env"
)

//Register your feature here
type Features struct {
	Route http.Handler
	Env   *env.Environment
} 

func Init() *Features {
	var ft Features
	ft.Route = route.Init()
	ft.Env   = env.Init()
	return &ft
}
