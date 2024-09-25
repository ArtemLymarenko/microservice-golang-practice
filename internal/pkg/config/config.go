package commonconfig

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Env string

const (
	EnvLocal Env = "local"
	EnvDev   Env = "dev"
	EnvProd  Env = "prod"
)

func expandEnvVars(input string) string {
	return os.Expand(input, func(key string) string {
		return os.Getenv(key)
	})
}

func parseYAMLWithEnv[T any](file string) *T {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil
	}

	contentWithEnv := expandEnvVars(string(content))

	var config T
	err = yaml.Unmarshal([]byte(contentWithEnv), &config)
	if err != nil {
		return nil
	}

	return &config
}

func MustGet[T any](path string) *T {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env")
	}

	config := parseYAMLWithEnv[T](path)
	if config == nil {
		log.Fatal("error parsing yaml")
	}

	return config
}
