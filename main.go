package main

import (
	corehttp "net/http"
	"template/config"
	"template/http"
	"template/http/handler"
	"template/http/middleware"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(config.LoadConfiguration,
			http.NewHTTPServer,
			fx.Annotate(
				http.NewRouter,
				fx.ParamTags(`group:"middlewares"`, `group:"routes"`),
			),
			http.AsRoute(handler.NewEchoHandler),
			http.AsMiddleware(middleware.NewSimpleMiddleware),
			http.AsMiddleware(middleware.NewSecondMiddleware),
			zap.NewExample),
		fx.Invoke(func(*corehttp.Server) {}),
	).Run()
}
