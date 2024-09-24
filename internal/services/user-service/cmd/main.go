package main

import (
	"os"
	"project-management-system/internal/pkg/logger"
	postgresql "project-management-system/internal/pkg/storage/posgres"
	"project-management-system/internal/user-service/internal/config"
)

func main() {
	cfg := config.MustGet()
	logs := logger.Setup(cfg.Env)
	logs.Info("Hello World!")
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
