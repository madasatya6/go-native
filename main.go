package main

import (
    "fmt"
	"log"
    "sync"
    "net/http"
    "crypto/tls"

	"github.com/madasatya6/go-native/system"
)

/**
* GO CLEAN ARCHITECTURE
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

    //advanced server
	//server.TLSConfig = getTlsConfig() //uncomment to activate advanced server

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

func getTlsConfig() *tls.Config {
	certPair1, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalln("Failed to start web server", err)
	}

	tlsConfig := new(tls.Config)
	tlsConfig.NextProtos = []string{"http/1.1"}
	tlsConfig.MinVersion = tls.VersionTLS12
	tlsConfig.PreferServerCipherSuites = true

	tlsConfig.Certificates = []tls.Certificate{
		certPair1, /** add other certificates here **/
	}
	tlsConfig.BuildNameToCertificate()

	tlsConfig.ClientAuth = tls.VerifyClientCertIfGiven
	tlsConfig.CurvePreferences = []tls.CurveID{
		tls.CurveP521,
		tls.CurveP384,
		tls.CurveP256,
	}
	tlsConfig.CipherSuites = []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	}

	return tlsConfig
}

