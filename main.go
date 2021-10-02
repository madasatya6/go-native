package main

import (
    "fmt"
	"log"

	"github.com/madasatya6/go-native/system/feature"
)

func main() {
    fmt.Println("Golang Clean Architecture")
	
	//fmt.Printf("%T", router)
    router := feature.Init()
    fmt.Println("Connected to port 9090")
    log.Fatal(http.ListenAndServe(":9090", router))
}

