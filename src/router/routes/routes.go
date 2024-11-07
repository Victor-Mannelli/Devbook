package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	Auth     bool
}

func Config(router *mux.Router) *mux.Router {
	routes := loginRoutes

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("/css"))

	return router
}
