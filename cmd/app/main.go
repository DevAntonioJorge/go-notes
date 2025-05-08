package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main(){
	if err := godotenv.Load(); err != nil{
		log.Fatalf("Error loading env: %v", err)
	} 
	e := echo.New()
	cfg := GetConfig()
    
	MapRoutes(e, cfg.JWTSecret, cfg.Port)
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
