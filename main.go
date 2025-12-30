package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=postgres dbname=sensors sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("Не удалось подключиться к БД: %v\n", err)
	}
	fmt.Println("Подключение к БД успешно!")
}
