package config

import (
	"project-management-system/internal/pkg/config"
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

func MustGet() *Config {
	return commonconfig.MustGet[Config]()
}
