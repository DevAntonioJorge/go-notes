package main

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MapRoutes(e *echo.Echo, secret string, port string, userHandler *UserHandler) {

	api := e.Group("/api")
	api.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			Skipper: middleware.DefaultSkipper,
			AllowOrigins: []string{"http://localhost" + port},
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
	metrics := api.Group("/metrics")
	metrics.Use(echoprometheus.NewMiddleware("api"),)
	metrics.GET("/", echoprometheus.NewHandler())
	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	users := api.Group("/users")

	public := users.Group("/public")
	{
		public.POST("/register", userHandler.RegisterHandler)
		public.POST("/login", userHandler.LoginHandler)
	}
	
	protected := users.Group("/protected")
	{
		protected.PATCH("/update-password", userHandler.UpdatePasswordHandler)
	}

	protected.Use(echojwt.WithConfig(echojwt.Config{
		Skipper: middleware.DefaultSkipper,
		SigningKey: []byte(secret),
		SigningMethod: "HS256",
		ContextKey: "user_id",
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			token, err := jwt.Parse(auth, func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256"{
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Sign method invalid")
				}
				return []byte(secret), nil
			})
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid Token")
			}
			return token, nil
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or absent token")
		},
		/*	
			Exemplo de uso
			user := c.Get("user_id").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
		*/
	}))
	

}
