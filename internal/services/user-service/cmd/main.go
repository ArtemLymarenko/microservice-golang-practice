package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"project-management-system/internal/pkg/storage"
	"project-management-system/internal/user-service/internal/app"
	"project-management-system/internal/user-service/internal/config"
	v1 "project-management-system/internal/user-service/internal/interfaces/rest/v1"
)

func main() {
	cfg := config.MustGet()

	postgres, err := storage.NewPostgres(cfg.Postgres, cfg.Env)
	if err != nil {
		logrus.Fatal(err.Error())
		os.Exit(1)
	}

	handlers, err := v1.InitializeHandlers(postgres, cfg.Service.Timeout)
	if err != nil {
		logrus.Fatal(err.Error())
		os.Exit(1)
	}

	router := v1.InitializeRouter(handlers)

	path, err := app.BuildHttpPath(cfg.HttpServer.Addr, cfg.HttpServer.Port)
	if err != nil {
		logrus.Fatal(err.Error())
		os.Exit(1)
	}

	server := &http.Server{
		Addr:         path,
		Handler:      router,
		ReadTimeout:  cfg.HttpServer.Timeout,
		WriteTimeout: cfg.HttpServer.Timeout,
		IdleTimeout:  cfg.HttpServer.IdleTimeout,
	}

	application := app.New(postgres, server)
	application.Start()
}
