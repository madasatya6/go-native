package routes 

import (
	"net/http"
	"github.com/gorilla/mux"
)

func WebRoute(route *mux.Router) *mux.Router {
	route.HandleFunc("/", index).Methods("GET")
	route.HandleFunc("/mux", intro).Methods("GET")
	return route
}

func index(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hallo World!"))
}

func intro(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla Mux!"))
}