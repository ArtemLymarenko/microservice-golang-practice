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
	JWT        JWT              `yaml:"jwt"`
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

type JWT struct {
	Secret     string        `yaml:"secret"`
	AccessExp  time.Duration `yaml:"accessExp"`
	RefreshExp time.Duration `yaml:"refreshExp"`
}

type App struct {
	CodeName string `yaml:"codeName"`
}

func requireConfigPath(env commonconfig.Env) string {
	switch env {
	case commonconfig.EnvLocal:
		return "resources/config/local.yaml"
	case commonconfig.EnvDev:
		return "resources/config/local.yaml"
	case commonconfig.EnvProd:
		return "resources/config/local.yaml"
	default:
		return ""
	}
}

func MustGet() *Config {
	configPath := requireConfigPath(commonconfig.EnvLocal)
	return commonconfig.MustGet[Config](configPath)
}
