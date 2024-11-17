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
	{
		URI:      "/posts/{postId}/like",
		Method:   http.MethodPost,
		Function: controllers.LikePost,
		Auth:     true,
	},
	{
		URI:      "/posts/{postId}/dislike",
		Method:   http.MethodPost,
		Function: controllers.DislikePost,
		Auth:     true,
	},
}
