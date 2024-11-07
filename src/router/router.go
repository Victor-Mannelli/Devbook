package router

import (
	"devbook/src/router/routes"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
