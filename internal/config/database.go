package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectDB(url string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	return conn
}

func ConnectMongoDB(url string) *mongo.Client {
	client, err := mongo.Connect(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB!")
	return client
}
