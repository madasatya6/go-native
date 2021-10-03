package route

import (
	"github.com/gorilla/mux"

	"github.com/madasatya6/go-native/routes"
)

type SystemRoute interface{
	WebRoute(route *mux.Router)
	APIRoute(route *mux.Router)
}

type typeRoute struct {
	Web *mux.Router
	API *mux.Router
}

func (t *typeRoute) WebRoute(route *mux.Router) {
	t.Web = routes.WebRoute(route)
}

func (t *typeRoute) APIRoute(route *mux.Router) {
	api := route.PathPrefix("/api").Subrouter()
	t.API = routes.APIRoute(api)
}

func Init() *mux.Router {

	var kind typeRoute
	var system SystemRoute

	route := mux.NewRouter()
	system = &kind

	system.WebRoute(route)
	system.APIRoute(route)
	return route
}
