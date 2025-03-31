package main

import (
	"fmt"
	"sync"
	"time"

	"math/rand"
)

type Task struct {
	ID int
}

func Worker(id int, tasks <-chan Task, semaphore chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		semaphore <- struct{}{}

		start := time.Now()
		fmt.Printf("Worker %d processing task %d\n", id, task.ID)
		time.Sleep(time.Duration(rand.Intn(2)+1) * time.Second) // Симуляция работы
		fmt.Printf("Worker %d finished task %d. Time: %d \n", id, task.ID, time.Since(start))

		<-semaphore
	}
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	const numWorkers = 3 // Количество горутин-воркеров
	const numTasks = 10  // Количество задач
	const maxConcurrent = 2

	tasks := make(chan Task, numTasks)
	semaphore := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup

	// Запускаем воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(i, tasks, semaphore, &wg)
	}

	// Отправляем задачи в канал
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i}
		fmt.Printf("Task %d received\n", i)
	}

	close(tasks) // Закрываем канал после отправки всех задач
	wg.Wait()    // Ждем завершения всех воркеров

	fmt.Println("All tasks processed")
}
