package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/andrew-nino/atm_v1/config"
	handler "github.com/andrew-nino/atm_v1/internal/controller"
	"github.com/andrew-nino/atm_v1/internal/service"
	"github.com/andrew-nino/atm_v1/pkg/server"

	log "github.com/sirupsen/logrus"
)

// Initialization and start of critical components.
func Run(configPath string) {

	// Configuration
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Logger
	SetLogrus(cfg.Log.Level)

	// Imitation repository
	repository := make(map[int]*service.Account)

	// Services dependencies
	log.Info("Initializing services...")
	service := service.NewService(repository)
	handlers := handler.NewHandler(service)

	// HTTP server
	log.Info("Starting http server...")
	srv := new(server.Server)

	go func() {
		if err := srv.Run(cfg.HTTP.Port, handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Print("App " + cfg.App.Name + " version: " + cfg.App.Version + " Started")

	// Waiting signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print(cfg.App.Name + " Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
