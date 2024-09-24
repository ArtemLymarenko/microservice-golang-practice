package main

import (
	"fmt"
	"project-management-system/internal/pkg/logger"
	"project-management-system/internal/services/project-service/internal/config"
)

func main() {
	cfg := config.New()
	logs := logger.Setup(cfg.Env)
	fmt.Println(cfg)
	fmt.Println(logs)
}
