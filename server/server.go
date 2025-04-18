package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bunkai888/hello-microservice/config"
	"github.com/bunkai888/hello-microservice/module/middleware/middlewareHandler"
	"github.com/bunkai888/hello-microservice/module/middleware/middlewareRepository"
	"github.com/bunkai888/hello-microservice/module/middleware/middlewareUsecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	server struct {
		app        *echo.Echo
		db         *mongo.Client
		cfg        *config.Config
		middleware middlewareHandler.MiddlewareHandlerService
	}
)

func newMiddleware(cfg *config.Config) middlewareHandler.MiddlewareHandlerService {
	repo := middlewareRepository.NewMiddlewareRepository()
	usecase := middlewareUsecase.NewMiddlewareUsecase(repo)
	return middlewareHandler.NewMiddlewareHandler(cfg, usecase)
}

func (s *server) gracefulShutdow(pctx context.Context, quit <-chan os.Signal) {
	log.Printf("Start service: %s", s.cfg.App.Name)
	<-quit
	log.Printf("Shutting down service: %s", s.cfg.App.Name)

	ctx, cancle := context.WithTimeout(pctx, 10*time.Second)
	defer cancle()

	if err := s.app.Shutdown(ctx); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func (s *server) httpListening() {
	if err := s.app.Start(s.cfg.App.Url); err != http.ErrServerClosed {
		log.Fatalf("Error Listening: %v", err)
	}
}

func Start(pctx context.Context, cfg *config.Config, db *mongo.Client) {
	s := &server{
		app:        echo.New(),
		db:         db,
		cfg:        cfg,
		middleware: newMiddleware(cfg),
	}

	//Basic Middleware
	//Request Timeout
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Error: Request Timeout",
		Timeout:      30 * time.Second,
	}))

	//CORS

	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
	}))

	//Body Limit
	s.app.Use(middleware.BodyLimit("10M"))

	//Custom Middleware

	switch s.cfg.App.Name {
	case "auth":
		s.authService()
	case "player":
		s.playerService()
	case "item":
		s.itemService()
	case "inventory":
		s.InventoryService()
	case "payment":
		s.PaymentService()
	}

	//Graceful Shutdow
	quite := make(chan os.Signal, 1)
	signal.Notify(quite, syscall.SIGINT, syscall.SIGTERM)

	s.app.Use(middleware.Logger())

	go s.gracefulShutdow(pctx, quite)

	//httpListening
	s.httpListening()
}
