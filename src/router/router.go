package router

import (
	"dev-book/src/router/routes"

	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	router := mux.NewRouter()
	return routes.SetUpRoutes(router)
}
