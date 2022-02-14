package routes

import (
	"dev-book/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.GetUsers,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodGet,
		Function:              controllers.GetUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodPatch,
		Function:              controllers.UpdateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequireAuthentication: false,
	},
}
