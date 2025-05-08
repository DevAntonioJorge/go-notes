package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

func main(){
	if err := godotenv.Load(); err != nil{
		log.Fatalf("Error loading env: %v", err)
	} 
	e := echo.New()
	cfg := GetConfig()
    
	MapRoutes(e, cfg.JWTSecret)
	go AddMetrics(cfg.MetricsPort)
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

func AddMetrics(port string) {
	metrics := echo.New()
	metrics.GET("/", echoprometheus.NewHandler())
	if err := metrics.Start(port); err != nil && !errors.Is(err, http.ErrServerClosed){
		log.Fatal(err)
	}
}