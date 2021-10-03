package routes 

import (
	"net/http"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", index).Methods("GET")
	return route
}

func index(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hallo World!"))
}