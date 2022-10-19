package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	conn, err := sql.Open("pgx", "postgres://postgres:mysecretpassword@192.168.59.106:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	if err := conn.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	cancel()
}
