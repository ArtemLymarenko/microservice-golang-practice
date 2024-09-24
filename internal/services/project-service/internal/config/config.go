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

type Postgres struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Dialect  string `yaml:"dialect"`
	Port     int    `yaml:"port"`
	PoolMin  int    `yaml:"poolMin"`
	PoolMax  int    `yaml:"poolMax"`
}

type HttpServer struct {
	Addr        string        `yaml:"addr"`
	Port        int           `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idleTimeout"`
}

func New() *Config {
	return commonconfig.MustGet[Config]("configs/configs.yaml")
}
