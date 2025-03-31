package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	connStr := "postgres://user:password@localhost:5432/mydb?sslmode=disable"

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("БД недоступна:", err)
	}

	fmt.Println("Успешное подключение к PostgreSQL!")

	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal("Ошибка выполнения запроса:", err)
	}
	defer rows.Close()

	sql := "Select id from users"
	_, err = db.Exec(sql, nil)
}
