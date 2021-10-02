package main

import (
    "fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func main() {
    fmt.Println("Golang Clean Architecture")

	router := mux.NewRouter()
    router.HandleFunc("/", index).Methods("GET")
    fmt.Println("Connected to port 9090")
    log.Fatal(http.ListenAndServe(":9090", router))
}

func index(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hallo World!"))
}