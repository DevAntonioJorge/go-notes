package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main(){
	if err := godotenv.Load(); err != nil{
		log.Fatalf("Error loading env: %v", err)
	} 
	e := echo.New()
	cfg := GetConfig()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"message": "OK"})
	})

	e.Logger.Fatal(e.Start(cfg.Port))
}