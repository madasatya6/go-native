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
	utils.SetContext(w,r)
	var data = map[string]interface{}{
		"title": "Learning html/template Actions",
	}
    utils.Render(w, "example/form-validation.html", data)
}

func Validate(w http.ResponseWriter, r *http.Request) {

	var form Form
	form.Nama = r.PostFormValue("nama")
	form.Alamat = r.PostFormValue("alamat")
	form.Umur = r.PostFormValue("umur")

	if err := validation.FormErrorID(w, r, form); err != nil {
		http.Redirect(w,r, "/example/form", http.StatusSeeOther)
	}

	var data = map[string]interface{}{
		"title": "Learning html/template Actions",
		"content": form,
	}

	utils.Render(w, "example/form-result.html", data)
}

