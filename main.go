package main

import (
    "fmt"
	"log"
    "sync"
    "net/http"

	"github.com/madasatya6/go-native/system"
)

/**
* CLEAN ARCHITECTURE GOLANG
* 
* @author Mada Satya Bayu Ambika
* @version 1.13
* @link https://github.com/madasatya6
* @gorillamux
*
* @access public
* @note change "madasatya6/go-native" according to the project name in all .go files
*/

func main() {
    var wg sync.WaitGroup
    fmt.Println("Go Clean Architecture by madasatya6")
	
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

