package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:      "/signup",
		Method:   http.MethodGet,
		Function: controllers.SingUpPage,
		Auth:     false,
	},
}
