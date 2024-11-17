package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:      "/posts",
		Method:   http.MethodPost,
		Function: controllers.CreatePost,
		Auth:     true,
	},
}
