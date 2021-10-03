package route

import (
	"net/http"
	
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
	return route
}
