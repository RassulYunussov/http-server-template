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

func NewPublicHTTPServer(p PublicHttpServerPamameters) *PublicServer {
	return &PublicServer{create(p.Lifecycle, p.Log, p.Config.PublicHttpServer, p.Router)}
}

type PublicHttpServerPamameters struct {
	fx.In
	Lifecycle fx.Lifecycle
	Log       *zap.Logger
	Config    config.Configuration
	Router    *muxtrace.Router `name:"public-router"`
}

var PublicServerModule = fx.Module("public-server",
	fx.Provide(
		NewPublicHTTPServer,
		router.NewPublicRouter,
		router.AsRoute(handler.NewEchoHandler, "public-routes"),
		middleware.AsMiddleware(middleware.NewSimpleMiddleware, "public-middlewares"),
		middleware.AsMiddleware(middleware.NewSecondMiddleware, "public-middlewares"),
	),
	fx.Invoke(func(*PublicServer) {}),
)
