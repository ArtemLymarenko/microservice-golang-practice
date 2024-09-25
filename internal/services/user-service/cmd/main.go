package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"project-management-system/internal/pkg/storage"
	"project-management-system/internal/user-service/internal/config"
)

func main() {
	cfg := config.MustGet()
	s, err := storage.NewPostgres(cfg.Postgres, cfg.Env)
	if err != nil {
		logrus.Fatal(err.Error())
		os.Exit(1)
	}
	err = s.CloseConnection()
	if err != nil {
		logrus.Info(err.Error())
		os.Exit(1)
	}
}
