package xdsconfig

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	CertServerPort         int `yaml:"cert_server_port" env-required:"true"`
	RouteServerPort        int `yaml:"route_server_port" env-required:"true"`
	EndpointManagementPort int `yaml:"endpoint_management_port" env-required:"true"`
	XdsPort                int `yaml:"xds_port" env-required:"true"`
}

var once sync.Once

func New(configPath string) (*Config, error) {
	var err error
	cfg := Config{}
	once.Do(func() {
		err = cleanenv.ReadConfig(configPath, &cfg)
	})

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
