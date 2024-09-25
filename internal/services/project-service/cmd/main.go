package main

import (
	"os"
	"project-management-system/internal/pkg/logger"
	postgresql "project-management-system/internal/pkg/storage"
	"project-management-system/internal/project-service/internal/config"
)

func main() {
	cfg := config.New()
	logs := logger.Setup(cfg.Env)

	s, err := postgresql.New(cfg.Postgres, cfg.Env, logs)
	if err != nil {
		logs.Info(err.Error())
		os.Exit(1)
	}
	err = s.CloseConnection()
	if err != nil {
		logs.Info(err.Error())
		os.Exit(1)
	}
}
