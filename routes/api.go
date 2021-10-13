package routes 

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func APIRoute(route *mux.Router) *mux.Router {
	var apikey = route.PathPrefix("/key").Subrouter()
	route.HandleFunc("/mux", apimux).Methods("GET")
	apikey.HandleFunc("/index", apikeymenu).Methods("GET")
	return route
}

func apimux(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status": true,
		"message": "Hallo Gorilla Mux!",
	}
    json.NewEncoder(w).Encode(response)
}

func apikeymenu(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status": true,
		"message": "Page with API KEY!",
	}
    json.NewEncoder(w).Encode(response)
}