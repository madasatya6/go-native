package system

import (

	"github.com/madasatya6/go-native/routes"
)

type Features struct {
	Route routes.Init
} 

func Init() *Features {
	var ft Features
	return &Features
}
