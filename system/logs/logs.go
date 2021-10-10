package logs

import (
	"net/http"
	"log"
	"os"

	"github.com/gorilla/mux"
	clog "github.com/madasatya6/go-native/applications/middleware/logger"
	klog "github.com/go-kit/kit/log"
)

type Log struct{}

func (t *Log) SetRouter(router *mux.Router) {
	var logger klog.Logger
	// Logfmt is a structured, key=val logging format that is easy to read and parse
	logger = klog.NewLogfmtLogger(klog.NewSyncWriter(os.Stderr))
	// Direct any attempts to use Go's log package to our structured logger
	log.SetOutput(klog.NewStdlibAdapter(logger))
	// Log the timestamp (in UTC) and the callsite (file + line number) of the logging
	// call for debugging in the future.
	logger = klog.With(logger, "ts", klog.DefaultTimestampUTC, "loc", klog.DefaultCaller)
	// Create an instance of our LoggingMiddleware with our configured logger
	loggingMiddleware := clog.LoggingMiddleware(logger)
	loggingMiddleware(router)
		
	router.Use(loggingMiddleware)
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