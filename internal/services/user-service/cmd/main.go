package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"project-management-system/internal/pkg/storage"
	"project-management-system/internal/user-service/internal/config"
	"project-management-system/internal/user-service/internal/domain/model"
	"project-management-system/internal/user-service/internal/domain/repository/postgres"
)

func main() {
	cfg := config.MustGet()
	s, err := storage.NewPostgres(cfg.Postgres, cfg.Env)
	if err != nil {
		logrus.Fatal(err.Error())
		os.Exit(1)
	}

	connection, _ := s.GetConnection()
	ur := postgres.NewUsersRepository(connection, postgres.NewUserInfoRepository(connection))

	ctx := context.Background()
	user, _ := model.NewUser("a3@a.com", "pass", "first", "last")
	err = ur.Save(ctx, user)
	if err != nil {
		logrus.Info(err.Error())
		os.Exit(1)
	}
	found, err := ur.FindByIdWithInfo(ctx, user.Id)
	if err != nil {
		logrus.Info(err.Error())
		os.Exit(1)
	}
	fmt.Println(found)

	err = s.CloseConnection()
	if err != nil {
		logrus.Info(err.Error())
		os.Exit(1)
	}
}
