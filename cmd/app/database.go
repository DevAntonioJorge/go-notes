package main

import (
	"context"
	"log"
	"os"
	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer conn.Close(context.Background())
	return conn
}