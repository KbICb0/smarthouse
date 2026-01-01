package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	// Строка подключения
	connStr := "postgres://postgres:postgres@localhost:5432/sensors?sslmode=disable"

	// Подключение
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		fmt.Printf("Не удалось подключиться: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	// Проверка
	err = conn.Ping(ctx)
	if err != nil {
		fmt.Printf("Ошибка ping: %v\n", err)
		return
	}

	fmt.Println("Подключение успешно!")

	rows, err := conn.Query(context.Background(), "select * from rooms")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var room_name string
		var temp_lo int
		var temp_hi int

		err := rows.Scan(&room_name, &temp_lo, &temp_hi)
		if err != nil {
			fmt.Printf("Ошибка чтения: %s\n", err)
			continue
		}
		fmt.Printf("Температура в комнате #%s с  %d до %d\n", room_name, temp_lo, temp_hi)
	}
}
