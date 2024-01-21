package config

import (
	"github.com/go-yaml/yaml"
	"io"
	"log"
	"os"
)

type Config struct {
	GrpcServerHost string `yaml:"grpc_server_host"`
	GrpcServerPort uint   `yaml:"grpc_server_port"`
}

func MustRun() *Config {
	config := new(Config)

	file, err := os.Open("./config/local.yaml")
	if err != nil {
		log.Fatalf("%v", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("%v", err)
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		log.Fatalf("%v", err)
	}

	return config
}
