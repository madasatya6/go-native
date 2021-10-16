package example

import (
	"fmt"
	"net/http"

	"github.com/madasatya6/go-native/helpers/utils"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla Mux!"))
}

func Session(w http.ResponseWriter, r *http.Request) {

	session, _ := utils.Session(r)
	session.Values["name"] = "Gorilla Mux"
	err := session.Save(r,w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := fmt.Sprintf("Tutorial Session %v ", session.Values["name"])
    w.Write([]byte(data))
}

func FlashSession(w http.ResponseWriter, r *http.Request) {

	utils.SetFlashdata(w, r, "Name", "Mada Satya")

	data := fmt.Sprintf("Tutorial Session %v ", utils.GetFlashdata(w, r, "Name"))
    w.Write([]byte(data))
}

