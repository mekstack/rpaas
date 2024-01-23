package config

import (
	"github.com/go-yaml/yaml"
	"io"
	"log"
	"os"
)

type GrpcServer struct {
	Host string `yaml:"host"`
	Port uint   `yaml:"port"`
}

type Redis struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
}

type Config struct {
	GrpcServer GrpcServer `yaml:"grpc_server"`
	Redis      Redis      `yaml:"redis"`
}

func MustConfig() *Config {
	config := new(Config)

	file, err := os.Open(os.Getenv("CONFIG_PATH"))
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
