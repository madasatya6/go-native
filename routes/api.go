package routes 

import (
	"net/http"
	"github.com/gorilla/mux"
)

func APIRoute(route *mux.Router) *mux.Router {
	route.HandleFunc("/mux", apimux).Methods("GET")
	return route
}

func apimux(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Routing API mux !"))
}
