package main

import (
	"fmt"
	"net/http"
	"project_1/pkg/config"
	"project_1/pkg/middleware"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	err := config.InitDB()
	if err != nil {
		return
	}

	fmt.Println("DB initialized and running")

	// ROUTER UNTIL NOW
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)

	server := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("Server started on :8080")
	server.ListenAndServe()

}
