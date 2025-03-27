package main

import (
	"context"
	"day_1/middlewares"
	"day_1/routes"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := routes.SetupRoutes()
	server := &http.Server{
		Addr:    ":8080",
		Handler: middlewares.Logger(middlewares.CORS(router)),
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		fmt.Println("Server started on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server error:", err)
		}
	}()

	<-stop
	fmt.Println("Server stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Server shutdown error:", err)
	}

}
