package routes 

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func APIRoute(route *mux.Router) *mux.Router {
	route.HandleFunc("/mux", apimux).Methods("GET")
	return route
}

func apimux(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status": true,
		"message": "Hallo Gorilla Mux!",
	}
    json.NewEncoder(w).Encode(response)
}
