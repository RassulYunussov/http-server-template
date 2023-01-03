package config

import (
	"go.uber.org/config"
)

type Configuration struct {
	Application struct {
		Name string
	}
	Httpserver struct {
		Port           int16
		Readtimeout    int
		Writetimeout   int
		Requesttimeout int
	}
	Database struct {
		Username string
		Password string
	}
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
