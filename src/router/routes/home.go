package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var homeRoute = Route{
	URI:      "/home",
	Method:   http.MethodGet,
	Function: controllers.HomePage,
	Auth:     true,
}
