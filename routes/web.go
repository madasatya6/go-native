package routes 

import (
	"github.com/gorilla/mux"
	"github.com/madasatya6/go-native/applications/controllers/example"
)

func WebRoute(route *mux.Router) *mux.Router {
	route.HandleFunc("/", example.Welcome).Methods("GET")
	route.HandleFunc("/example/session", example.Session).Methods("GET")

	return route
}



