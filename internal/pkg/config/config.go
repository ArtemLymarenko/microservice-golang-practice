package commonconfig

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

type Env string

const (
	EnvLocal Env = "local"
	EnvDev   Env = "dev"
	EnvProd  Env = "prod"
)

func MustGet[T any](path string) *T {
	if strings.Trim(path, " ") == "" {
		log.Fatal("config path was not found")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("path %v does not exists", path)
	}

	file, err := os.Open(path)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalf("error closing file: %v", err)
		}
	}()

	if err != nil {
		log.Fatalf("cannot open the file")
	}

	var config T
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("error parsing the file")
	}

	return &config
}
