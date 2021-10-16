package example

import (
	"net/http"

	"github.com/madasatya6/go-native/helpers/utils"
)

func FormValidation(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{}{
		"title": "Learning html/template Actions",
	}
    utils.Render(w, "example/form-validation.html", data)
}

func Validate(w http.ResponseWriter, r *http.Request) {
	
}

