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
		Auth:     true,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodGet,
		Function: controllers.UserPage,
		Auth:     true,
	},
	{
		URI:      "/users/{userId}/unfollow",
		Method:   http.MethodPost,
		Function: controllers.Unfollow,
		Auth:     true,
	},
	{
		URI:      "/users/{userId}/follow",
		Method:   http.MethodPost,
		Function: controllers.Follow,
		Auth:     true,
	},
}
