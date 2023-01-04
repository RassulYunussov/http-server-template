package main

import (
	"template/config"
	"template/http-server"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(zap.NewExample),
		config.ConfigurationModule,
		http.HttpServersModule,
	).Run()
}
