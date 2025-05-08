package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MapRoutes(e *echo.Echo) {

	api := e.Group("/api")
	api.Use(
		middleware.CORS(),
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
}
