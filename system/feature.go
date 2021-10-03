package system

import (
	"net/http"
	"github.com/madasatya6/go-native/system/route"
)

//Register your feature here
type Features struct {
	Route http.Handler
} 

func Init() *Features {
	var ft Features
	ft.Route = route.Init()
	return &ft
}
