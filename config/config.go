package config

import (
	"go.uber.org/config"
	"go.uber.org/fx"
)

type Configuration struct {
	Application struct {
		Name string
	}
	PrivateHttpServer HttpServer `yaml:"privateHttpServer"`
	PublicHttpServer  HttpServer `yaml:"publicHttpServer"`
	Database          struct {
		Username string
		Password string
	}
}

type HttpServer struct {
	Port           int16
	ReadTimeout    int `yaml:"readTimeout"`
	WriteTimeout   int `yaml:"writeTimeout"`
	RequestTimeout int `yaml:"requestTimeout"`
}

func LoadConfiguration() (Configuration, error) {
	var c Configuration
	cfg, err := config.NewYAML(config.File("config.yaml"))
	if err != nil {
		return c, err
	}

	if err := cfg.Get("").Populate(&c); err != nil {
		return c, err
	}

	return c, nil
}

var ConfigurationModule = fx.Module("application-configuration", fx.Provide(LoadConfiguration))
