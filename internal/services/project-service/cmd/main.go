package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	appUtil "project-management-system/internal/pkg/app"
	"project-management-system/internal/pkg/storage"
	"project-management-system/internal/project-service/internal/app"
	"project-management-system/internal/project-service/internal/config"
	v1 "project-management-system/internal/project-service/internal/interface/rest/v1"
)

func main() {
	cfg := config.New()
	postgres, err := storage.NewPostgres(cfg.Postgres, cfg.Env)
	if err != nil {
		logrus.Info(err.Error())
		os.Exit(1)
	}

	router := v1.MustGetGinRouter(postgres, cfg)
	path, err := appUtil.BuildHttpPath(cfg.HttpServer.Addr, cfg.HttpServer.Port)
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
