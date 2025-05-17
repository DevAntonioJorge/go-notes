package main

import (
	"context"
	"log"
	"os"

	"github.com/DevAntonioJorge/go-notes/internal/api"
	"github.com/DevAntonioJorge/go-notes/internal/config"
	"github.com/DevAntonioJorge/go-notes/internal/handlers"
	"github.com/DevAntonioJorge/go-notes/internal/repository"
	"github.com/DevAntonioJorge/go-notes/internal/service"
	"github.com/DevAntonioJorge/go-notes/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading env: %v", err)
	}
	cfg := config.GetConfig()
	db := config.ConnectDB(cfg.DBUrl)
	mgDB := config.ConnectMongoDB(cfg.MongoDBUrl)
	defer func() {
		if err := mgDB.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting to Mongo client")
		}
	}()

	logger := logger.New(os.Stdout, "", log.LstdFlags, logger.LevelInfo)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	folderRepository := repository.NewFolderRepository(mgDB.Database("go-notes"))
	folderService := service.NewFolderService(folderRepository)
	folderHandler := handlers.NewFolderHandler(folderService)
	server := api.NewServer(cfg.Port, cfg.JWTSecret, userHandler, folderHandler, logger)
	server.MapRoutes()

	if err := server.Run(); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
