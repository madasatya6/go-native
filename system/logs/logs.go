package logs

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

type Log struct{}

func (t *Log) SetRouter(router *mux.Router) {
	router.Use(LoggingMiddleware)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func Init() *Log {
	var lg Log
	return &lg 
}