package routes

import (
	"devbook/src/middlewares"
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
	routes = append(routes, usersRoutes...)
	routes = append(routes, postsRoutes...)
	routes = append(routes, logoutRoute)
	routes = append(routes, homeRoute)

	for _, route := range routes {
		if route.Auth {
			router.HandleFunc(route.URI,
				middlewares.Logger(middlewares.AuthExists(route.Function)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
