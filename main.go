package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func connectBD() (*pgx.Conn, error) { //для теста подключения, потом надо удалить
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
		return nil, err
	}

	fmt.Println("Подключение успешно!")
	return pgx.Connect(context.Background(), connStr)
}

func selectBdRooms() {
	conn, err := connectBD()
	if err != nil {
		fmt.Printf("Ошибка подключения: %v\n", err)
		return
	}
	defer conn.Close(context.Background())
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
		fmt.Println("Введите 1 для получения температуры в спальне, 2 для выхода:")

		fmt.Scan(&homeOption)
		switch homeOption {
		case 1:
			selectBdRooms()
		case 2:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Нет опции")
		}
	}

}
