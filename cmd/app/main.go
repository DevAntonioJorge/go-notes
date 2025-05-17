package main

import (
	"log"

	"github.com/DevAntonioJorge/go-notes/internal/api"
	"github.com/DevAntonioJorge/go-notes/internal/config"
	"github.com/DevAntonioJorge/go-notes/internal/handlers"
	"github.com/DevAntonioJorge/go-notes/internal/repository"
	"github.com/DevAntonioJorge/go-notes/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading env: %v", err)
	}
	cfg := config.GetConfig()
	db := config.ConnectDB(cfg.DBUrl)
	//mgDB := ConnectMongoDB(cfg.MongoDBUrl)
	/*defer func() {
		if err := mgDB.Disconnect(); err != nil{
			log.Fatalf("Error disconnecting to Mongo client")
		}
	}
	*/
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	server := api.NewServer(cfg.Port, cfg.JWTSecret, userHandler, nil)
	server.MapRoutes()

	if err := server.Run(); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
