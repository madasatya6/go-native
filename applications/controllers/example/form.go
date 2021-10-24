package example

import (
	"net/http"

	"github.com/madasatya6/go-native/helpers/utils"
	"github.com/madasatya6/go-native/helpers/packages/validation"
)

type Form struct{
	Nama	string	`json:"nama" form:"nama" validate:"required"` 
	Alamat	string	`json:"alamat" form:"alamat" validate:"required"`
	Umur	string	`json:"umur" form:"umur" validate:"required,numeric"`
}

func FormValidation(w http.ResponseWriter, r *http.Request) {

	var data = map[string]interface{}{
		"Response" : w,
		"Request" : r,
		"title": "Learning html/template Actions",
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1. 
	w.Header().Set("Pragma", "no-cache") // HTTP 1.0. 
	w.Header().Set("Expires", "0") // Proxies.

    utils.Render(w, "example/form-validation.html", data)
}

func Validate(w http.ResponseWriter, r *http.Request) {
	
	var form Form
	form.Nama = r.PostFormValue("nama")
	form.Alamat = r.PostFormValue("alamat")
	form.Umur = r.PostFormValue("umur")

	if err := validation.FormErrorID(w, r, form); err != nil {
		http.Redirect(w,r, "/example/form", http.StatusFound)
	}

	var data = map[string]interface{}{
		"title": "Learning html/template Actions",
		"content": form,
	}

	utils.Render(w, "example/form-result.html", data)
}

