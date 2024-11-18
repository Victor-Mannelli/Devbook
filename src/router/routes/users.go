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
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		Auth:     false,
	},
	{
		URI:      "/find-users",
		Method:   http.MethodGet,
		Function: controllers.FilteredUsersPage,
		Auth:     false,
	},
}
