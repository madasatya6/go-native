package boot

import (
	"net/http"
	"github.com/madasatya6/go-native/applications/config"
)

func Init(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		go config.RegisterBoots()
		next.ServeHTTP(w,r)
	})
}

