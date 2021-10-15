package example

import (
	"fmt"
	"net/http"

	"github.com/madasatya6/go-native/helpers/utils"
	"github.com/madasatya6/go-native/applications/config"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla Mux!"))
}

func Session(w http.ResponseWriter, r *http.Request) {
	session, err := utils.SessionStore.Get(r, config.SessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	session.Values["name"] = "Gorilla Mux"
	session.Save(r,w)

	data := fmt.Sprintf("Tutorial Session %v", session.Values["name"])

    w.Write([]byte(data))
}