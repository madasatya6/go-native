package main

import (
    "fmt"
	"log"
    "net/http"
	"github.com/madasatya6/go-native/system"
)

func main() {
    fmt.Println("Golang Clean Architecture")
	
	systems := system.Init()
    router := systems.Route
    fmt.Println("Connected to port 9090")
    log.Fatal(http.ListenAndServe(":9090", router))
}

