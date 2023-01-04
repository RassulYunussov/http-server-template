package http

import (
	"net/http"
	"template/config"
	"template/http-server/handler"
	"template/http-server/middleware"
	"template/http-server/router"

	"go.uber.org/fx"
	"go.uber.org/zap"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

type PublicServer struct {
	*http.Server
}

func NewPublicHTTPServer(lc fx.Lifecycle, log *zap.Logger, config config.Configuration, router *muxtrace.Router) *PublicServer {
	return &PublicServer{create(lc, log, config.PublicHttpServer, router)}
}

var PublicServerModule = fx.Module("public-server",
	fx.Provide(
		// public server
		fx.Annotate(
			NewPublicHTTPServer,
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
	),
	fx.Invoke(func(*PublicServer) {}),
)
