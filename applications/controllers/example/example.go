package example

import (
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla Mux!"))
}

func Session(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hallo World!"))
}