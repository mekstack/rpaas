package config

import (
	"fmt"
	"os"
	"strings"
)

var validEnvironments = []string{"Development", "Production"}

const (
	Development uint = iota
	Production
)

type GrpcServer struct {
	Addr string
}

type Redis struct {
	Addr string
}

type Config struct {
	GrpcServer  *GrpcServer
	Redis       *Redis
	Environment uint
}

func mustReadEnvironment() uint {
	value := readEnvOrSetDefault("NATAAS_ENVIRONMENT", "Development")
	envCode := Development
	switch value {
	case "Development":
		envCode = Development
	case "Production":
		envCode = Production
	default:
		panic(fmt.Errorf("Value %s is not valid for NATAAS_ENVIRONMENT, valid vales: "+strings.Join(validEnvironments, " "), value))
	}

	return envCode
}

func readEnvOrSetDefault(key string, def string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		os.Setenv(key, def)
		value = def
	}
	return value
}

func MustConfig() *Config {
	return &Config{
		Redis: &Redis{
			Addr: readEnvOrSetDefault("NATAAS_REDIS_ADDR", "127.0.0.1:6379"),
		},
		GrpcServer: &GrpcServer{
			Addr: readEnvOrSetDefault("NATAAS_GRPC_ADDR", "127.0.0.1:8080"),
		},
		Environment: mustReadEnvironment(),
	}
}
