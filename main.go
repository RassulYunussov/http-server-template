package main

import (
	"template/config"
	"template/http"
	"template/http/handler"
	"template/http/middleware"
	"template/http/router"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(config.LoadConfiguration,
			zap.NewExample,
			// public server
			fx.Annotate(
				http.NewPublicHTTPServer,
				fx.ParamTags(``, ``, ``, `name:"public-router"`),
			),
			fx.Annotate(
				router.NewRouter,
				fx.ParamTags(`group:"public-middlewares"`, `group:"public-routes"`),
				fx.ResultTags(`name:"public-router"`),
			),
			router.AsRoute(handler.NewEchoHandler, "public-routes"),
			middleware.AsMiddleware(middleware.NewSimpleMiddleware, "public-middlewares"),
			middleware.AsMiddleware(middleware.NewSecondMiddleware, "public-middlewares"),

			// private server
			fx.Annotate(
				http.NewPrivateHTTPServer,
				fx.ParamTags(``, ``, ``, `name:"private-router"`),
			),
			fx.Annotate(
				router.NewRouter,
				fx.ParamTags(`group:"private-middlewares"`, `group:"private-routes"`),
				fx.ResultTags(`name:"private-router"`),
			),
			router.AsRoute(handler.NewEchoHandler, "private-routes"),
			middleware.AsMiddleware(middleware.NewSimpleMiddleware, "private-middlewares"),
		),
		fx.Invoke(func(*http.PublicServer, *http.PrivateServer) {}),
	).Run()
}
