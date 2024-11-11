package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		URI:      "/",
		Method:   http.MethodGet,
		Function: controllers.LoginPage,
		Auth:     false,
	},
	{
		URI:      "/login",
		Method:   http.MethodGet,
		Function: controllers.LoginPage,
		Auth:     false,
	},
	{
		URI:      "/login",
		Method:   http.MethodPost,
		Function: controllers.Login,
		Auth:     false,
	},
}
