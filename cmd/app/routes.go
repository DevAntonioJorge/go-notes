package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo-jwt/v4"
)

func MapRoutes(e *echo.Echo, secret string) {

	api := e.Group("/api")
	api.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			Skipper: middleware.DefaultSkipper,
			AllowOrigins: []string{"http://localhost:8000"},
			AllowMethods: []string{http.MethodGet, http.MethodPatch, http.MethodDelete, http.MethodPost},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
			AllowCredentials: true,
		}),
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogLatency: true,
			LogStatus:  true,
			LogError:   true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				e.Logger.Debugf("REQUEST: uri=%v, status=%v, latency=%v, error=%v",
				v.URI, v.Status, v.Latency, v.Error)
				return nil
			},
		}),
		middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(30)),
		middleware.RequestID(),
		middleware.Recover(),
	)

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	users := api.Group("/users")
	//public := users.Group("/")
	protected := users.Group("/")
	
	protected.Use(echojwt.JWT([]byte(secret)))
}
