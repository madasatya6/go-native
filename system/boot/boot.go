package boot

import (
	"net/http"
	"github.com/madasatya6/go-native/applications/config"
	"github.com/madasatya6/go-native/system/times"
)

func Init(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		go config.RegisterBoots()
		go times.Init()
		next.ServeHTTP(w,r)
	})
}

