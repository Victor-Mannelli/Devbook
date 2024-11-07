package router

import "github.com/gorilla/mux"

func Setup() *mux.Router {
	return mux.NewRouter()
}
