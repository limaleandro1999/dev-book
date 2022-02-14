package main

import (
	"dev-book/src/config"
	"dev-book/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnvVars()

	fmt.Println(config.Port)
	fmt.Println("API running")

	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
