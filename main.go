package main

import (
    "fmt"
	"log"
    "sync"
    "net/http"

	"github.com/madasatya6/go-native/system"
)

func main() {
    var wg sync.WaitGroup
    fmt.Println("Golang Clean Architecture")
	
	systems := system.Init()
    router := systems.Route
    env := systems.Env
    systems.LogEntry.SetRouter(router)
    fmt.Println("Connected to port ", env.Port)

    lock := make(chan error)
    srv := &http.Server{
        Handler:      router, //type http.Handler
        Addr:         env.Port,
        WriteTimeout: env.WriteTimeout,
        ReadTimeout:  env.ReadTimeout,
    }

    wg.Add(1)
    go func(wg sync.WaitGroup){
        defer wg.Done()
        lock<- srv.ListenAndServe()
    }(wg)
    
    wg.Wait()

    log.Fatal(lock)
}

