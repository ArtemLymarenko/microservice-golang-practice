package logger

import (
	"log/slog"
	"os"
	"project-management-system/internal/pkg/config"
)

func Setup(env commonconfig.commonconfig) *slog.Logger {
	var log *slog.Logger

	switch env {
	case commonconfig.EnvLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		break
	case commonconfig.EnvDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		break
	case commonconfig.EnvProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		break
	}

	return log
}
