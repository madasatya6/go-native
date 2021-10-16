package utils

import (
	"net/http"
	"html/template"
)

func Render(w http.ResponseWriter, viewPath string, data map[string]interface{}){
	var location = "resource/views/"
	var tmpl = template.Must(template.New(location + viewPath).
		Funcs(FuncMap).
		ParseFiles(
			location + viewPath,
		))
		
	err := tmpl.ExecuteTemplate(w, "main", data)	
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func View(w http.ResponseWriter, viewPath string, data map[string]interface{}){
	var location = "resource/views/"
	var tmpl = template.Must(template.New("resource/views/backend/main-layout.html").
		Funcs(FuncMap).
		ParseFiles(
			"resource/views/backend/navbar.html",
			"resource/views/backend/sidebar.html",
			location + viewPath,
		))
		
	err := tmpl.ExecuteTemplate(w, "main", data)	
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}