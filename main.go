package main

import (
    "fmt"
	"log"
    "time"
    "sync"
    "net/http"

	"github.com/madasatya6/go-native/system"
)

func main() {
    var wg sync.WaitGroup
    fmt.Println("Golang Clean Architecture")
	
	systems := system.Init()
    router := systems.Route
    fmt.Println("Connected to port 9090")

    lock := make(chan error)
    srv := &http.Server{
        Handler:      router,
        Addr:         "127.0.0.1:9090",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    wg.Add(1)
    go func(wg sync.WaitGroup){
        defer wg.Done()
        lock<- srv.ListenAndServe()
    }(wg)
    
    wg.Wait()

    log.Fatal(lock)
}

