package main

import (
	"day_2/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes.RegisterBookRoutes(router)

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
