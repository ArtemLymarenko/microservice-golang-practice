package config

import (
	"project-management-system/internal/pkg/config"
	"time"
)

type Config struct {
	Env        commonconfig.Env `yaml:"env"`
	Postgres   Postgres         `yaml:"postgres"`
	HttpServer HttpServer       `yaml:"httpServer"`
	Service    Service          `yaml:"service"`
}

type HttpServer struct {
	Addr            string        `yaml:"addr"`
	Port            int           `yaml:"port"`
	Timeout         time.Duration `yaml:"timeout"`
	IdleTimeout     time.Duration `yaml:"idleTimeout"`
	ShutDownTimeout time.Duration `yaml:"shutDownTimeout"`
}

type Service struct {
	Timeout time.Duration `yaml:"timeout"`
}

func MustGet() *Config {
	const configPath = "resources/config/local.yaml"
	return commonconfig.MustGet[Config](configPath)
}
