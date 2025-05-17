package api

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DevAntonioJorge/go-notes/internal/handlers"
	"github.com/DevAntonioJorge/go-notes/pkg/logger"
	"github.com/labstack/echo/v4"
)

type Server struct {
	router        *echo.Echo
	port          string
	secret        string
	logger        *logger.Logger
	userHandler   *handlers.UserHandler
	folderHandler *handlers.FolderHandler
	//noteHandler *handlers.NoteHandler
}

func NewServer(port string, secret string, userHandler *handlers.UserHandler, folderHandler *handlers.FolderHandler, logger *logger.Logger) *Server {
	return &Server{
		router:        echo.New(),
		port:          port,
		secret:        secret,
		logger:        logger,
		userHandler:   userHandler,
		folderHandler: folderHandler,
	}
}

func (s *Server) Run() error {

	shutdown := make(chan error, 1)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		sig := <-quit

		s.logger.Debug("Signal captured: %v", sig.String())

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		shutdown <- s.router.Shutdown(ctx)
	}()

	err := s.router.Start(s.port)
	if err != nil {
		return err
	}

	if err = <-shutdown; err != nil {
		return err
	}
	return nil
}
