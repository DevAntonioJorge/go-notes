package main

import (
	"context"
	"log"
	"github.com/jackc/pgx/v5"
)

func ConnectDB(url string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer conn.Close(context.Background())
	return conn
}