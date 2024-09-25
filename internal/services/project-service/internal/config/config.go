package config

import (
	commonconfig "project-management-system/internal/pkg/config"
	"time"
)

type Config struct {
	Env        commonconfig.Env `yaml:"env"`
	Postgres   Postgres         `yaml:"postgres"`
	HttpServer HttpServer       `yaml:"httpServer"`
}

type HttpServer struct {
	Addr        string        `yaml:"addr"`
	Port        int           `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idleTimeout"`
}

func New() *Config {
	const configPath = "configs/configs.yaml"
	return commonconfig.MustGet[Config](configPath)
}
