package route

import (
	"github.com/gorilla/mux"

	"github.com/madasatya6/go-native/routes"
)

type SystemRoute interface{
	WebRoute()
	//APIRoute()
}

type typeRoute struct {
	Web *mux.Router
	//API *mux.Router
}

func (t *typeRoute) WebRoute(route *mux.Router) {
	t.Web = routes.WebRoute(route)
}

func Init() *mux.Router {
	var kind typeRoute
	route := mux.NewRouter()
	kind.WebRoute(route)
	return route
}
