package api

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-contrib/echoprometheus"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) MapRoutes() {

	api := s.router.Group("/api")
	api.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			Skipper:          middleware.DefaultSkipper,
			AllowOrigins:     []string{"http://localhost" + s.port},
			AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodDelete, http.MethodPost},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
			AllowCredentials: true,
		}),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Skipper: middleware.DefaultSkipper,
			Format:  `"${time_rfc3339} [${remote_ip}] "${method} ${uri}" ${status} ${latency} ${bytes_in} ${bytes_out}`,
		}),
		middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(30)),
		middleware.RequestID(),
		middleware.Recover(),
	)
	metrics := api.Group("/metrics")
	metrics.Use(echoprometheus.NewMiddleware("api"))
	metrics.GET("/", echoprometheus.NewHandler())
	api.GET("/", func(c echo.Context) error {
		s.logger.Info("Home")
		return c.String(http.StatusOK, "Hello, World!")
	})

	users := api.Group("/users")

	public := users.Group("/public")
	{
		public.POST("/register", s.userHandler.RegisterHandler)
		public.POST("/login", s.userHandler.LoginHandler)
	}

	protected := users.Group("/protected")
	{
		protected.PATCH("/update-password", s.userHandler.UpdatePasswordHandler)
	}

	protected.Use(echojwt.WithConfig(echojwt.Config{
		Skipper:       middleware.DefaultSkipper,
		SigningKey:    []byte(s.secret),
		SigningMethod: "HS256",
		ContextKey:    "user_id",
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			token, err := jwt.Parse(auth, func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					s.logger.Error("Sign method invalid: %v", t.Method.Alg())
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Sign method invalid")
				}
				return []byte(s.secret), nil
			})
			if err != nil {
				s.logger.Error("Error parsing token: %v", err)
				return nil, err
			}
			if !token.Valid {
				s.logger.Error("Invalid token")
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid Token")
			}
			return token, nil
		},
		ErrorHandler: func(c echo.Context, err error) error {
			s.logger.Error("Error parsing token: %v", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or absent token")
		},
		/*
			Exemplo de uso
			user := c.Get("user_id").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
		*/
	}))

}
