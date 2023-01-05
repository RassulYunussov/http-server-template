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

type PrivateServer struct {
	*http.Server
}

func NewPrivateHTTPServer(p PrivateHttpServerPamameters) *PrivateServer {
	return &PrivateServer{create(p.Lifecycle, p.Log, p.Config.PrivateHttpServer, p.Router)}
}

type PrivateHttpServerPamameters struct {
	fx.In
	Lifecycle fx.Lifecycle
	Log       *zap.Logger
	Config    config.Configuration
	Router    *muxtrace.Router `name:"private-router"`
}

var PrivateServerModule = fx.Module("private-server",
	fx.Provide(
		NewPrivateHTTPServer,
		router.NewPrivateRouter,
		router.AsRoute(handler.NewEchoHandler, "private-routes"),
		middleware.AsMiddleware(middleware.NewSimpleMiddleware, "private-middlewares"),
	),
	fx.Invoke(func(*PrivateServer) {}),
)
