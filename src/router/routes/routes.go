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
	routes := []Route{}

	for _, route := range routes {
		if route.Auth {
			r.HandleFunc(route.URI, route.Function).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, route.Function).Methods(route.Method)
		}
	}

	return r
}
