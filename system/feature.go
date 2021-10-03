package system

import (
	"net/http"
	"github.com/madasatya6/go-native/routes"
)

type Features struct {
	Route http.Handler
} 

func Init() *Features {
	var ft Features
	ft.Route = routes.Init()
	return &ft
}
