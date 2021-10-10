package route

import (
	"net/http"
	"html/template"
	
	"github.com/gorilla/mux"
	"github.com/madasatya6/go-native/routes"
)

type SystemRoute interface{
	StaticAsset(route *mux.Router)
	WebRoute(route *mux.Router)
	APIRoute(route *mux.Router)
}

type typeRoute struct {
	Web *mux.Router
	API *mux.Router
	Asset *mux.Route
}

func (t *typeRoute) WebRoute(route *mux.Router) {
	t.Web = routes.WebRoute(route)
}

func (t *typeRoute) APIRoute(route *mux.Router) {
	api := route.PathPrefix("/api").Subrouter()
	t.API = routes.APIRoute(api)
}

func (t *typeRoute) StaticAsset(route *mux.Router) {
	t.Asset = route.PathPrefix("/assets").Handler(http.FileServer(http.Dir("./resource/assets/")))
}

func Init() *mux.Router {

	var kind typeRoute
	var system SystemRoute

	route := mux.NewRouter()
	system = &kind

	system.StaticAsset(route)
	system.WebRoute(route)
	system.APIRoute(route)

	route.NotFoundHandler = http.HandlerFunc(NotFound)
	return route
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{}{}
	tmpl := template.Must(template.ParseFiles("resource/views/404.html"))
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
