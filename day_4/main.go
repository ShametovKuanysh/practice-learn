package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type key string

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Запрос обработан")
	case <-ctx.Done():
		http.Error(w, "Время обработки запроса истекло", http.StatusGatewayTimeout)
	}
}

func handlerWithValue(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), key("userID"), 12345)

	processRequest(ctx, w)
}

func processRequest(ctx context.Context, w http.ResponseWriter) {
	userID, ok := ctx.Value(key("userID")).(int)
	if !ok {
		http.Error(w, "Ошибка авторизации", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Обрабатываем запрос пользователя с ID: %d", userID)
}

// func worker(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Worker stopped")
// 			return
// 		default:
// 			fmt.Println("worker is working")
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

func main() {
	// deadline := time.Now().Add(3 * time.Second)
	// ctx, cancel := context.WithDeadline(context.Background(), deadline)
	// defer cancel()

	// go worker(ctx)

	// time.Sleep(4 * time.Second)
	// fmt.Println("Main function stopped")

	http.HandleFunc("/start", handler)
	http.HandleFunc("/2", handlerWithValue)
	fmt.Println("Сервер запущен на :8080")
	http.ListenAndServe(":8080", nil)
}
