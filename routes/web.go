package routes 

import (
	"net/http"
	"github.com/gorilla/mux"
)

func Init() *mux.RouterConnected {
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	return router
}

func index(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hallo World!"))
}