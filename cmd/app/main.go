package main

import (
	"context"
	"log"
	"os"

	"github.com/DevAntonioJorge/go-notes/internal/domain/repository"
	"github.com/DevAntonioJorge/go-notes/internal/domain/service"
	"github.com/DevAntonioJorge/go-notes/internal/infra/api"
	"github.com/DevAntonioJorge/go-notes/internal/infra/config"
	"github.com/DevAntonioJorge/go-notes/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	cfg := config.GetConfig()
	logger := logger.New(os.Stdout, "", log.LstdFlags, cfg.LogLevel)
	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading env: %v", err)
	}
	logger.Info("Env loaded")
	db := config.ConnectDB(cfg.DBUrl)
	logger.Info("DB connected")
	mgDB := config.ConnectMongoDB(cfg.MongoDBUrl)
	defer func() {
		if err := mgDB.Disconnect(context.Background()); err != nil {
			logger.Error("Error disconnecting to Mongo client")
		}
	}()
	logger.Info("MongoDB connected")

	repository := repository.NewRepository(db, mgDB)
	service := service.NewService(repository)
	server := api.NewServer(cfg.Port, cfg.JWTSecret, logger, service)
	server.MapRoutes()

	if err := server.Run(); err != nil {
		logger.Fatal("Error running server: %v", err)
	}
}
