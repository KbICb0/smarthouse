package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func connectbd() { //для теста подключения, потом надо удалить
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
}

func selectBdRooms() {
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

	rows, err := conn.Query(context.Background(), "select * from rooms")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var room_name string //Нужно сделать struct
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

func main() {
	var homeOption int
	for {
		fmt.Println("Введите 1 для теста подключения базы данных, 2 для получения температуры в спальне, 3 для выхода:")

		fmt.Scan(&homeOption)
		switch homeOption {
		case 1:
			connectbd()
		case 2:
			selectBdRooms()
		case 3:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Нет опции")
		}
	}

}
