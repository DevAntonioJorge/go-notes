package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DevAntonioJorge/go-notes/internal/repository"
	"github.com/DevAntonioJorge/go-notes/internal/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main(){
	if err := godotenv.Load(); err != nil{
		log.Fatalf("Error loading env: %v", err)
	} 
	e := echo.New()

	cfg := GetConfig()
    	db := ConnectDB(cfg.DBUrl)
	//mgDB := ConnectMongoDB(cfg.MongoDBUrl)
	/*defer func() {
		if err := mgDB.Disconnect(); err != nil{
			log.Fatalf("Error disconnecting to Mongo client")
		}
	}
	*/
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := NewUserHandler(userService)

	MapRoutes(e, cfg.JWTSecret, cfg.Port, userHandler)
	e.Logger.Fatal(Run(cfg.Port, e))
}

func Run(port string, e *echo.Echo) error{

	shutdown := make(chan error, 1)

	go func(){
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s:= <-quit

		e.Logger.Debugf("Signal captured: %v", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		shutdown <- e.Shutdown(ctx)
	}()
	
	err := e.Start(port)
	if err != nil{
		return err
	}

	if err = <- shutdown; err != nil {
		return err
	}
	return nil
}
