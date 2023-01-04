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

func NewPrivateHTTPServer(lc fx.Lifecycle, log *zap.Logger, config config.Configuration, router *muxtrace.Router) *PrivateServer {
	return &PrivateServer{create(lc, log, config.PrivateHttpServer, router)}
}

var PrivateServerModule = fx.Module("private-server",
	fx.Provide(
		fx.Annotate(
			NewPrivateHTTPServer,
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
	fx.Invoke(func(*PrivateServer) {}),
)
