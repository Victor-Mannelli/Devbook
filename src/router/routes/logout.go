package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var logoutRoute = Route{
	URI:      "/logout",
	Method:   http.MethodGet,
	Function: controllers.Logout,
	Auth:     true,
}
