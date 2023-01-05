package router

import (
	"template/http-server/middleware"

	"go.uber.org/fx"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

type PrivateRouterPamameters struct {
	fx.In
	Middlewares []middleware.Middleware `group:"private-middlewares"`
	Routes      []Route                 `group:"private-routes"`
}

type PrivateRouter struct {
	fx.Out
	Router *muxtrace.Router `name:"private-router"`
}

func NewPrivateRouter(p PrivateRouterPamameters) PrivateRouter {
	return PrivateRouter{Router: createRouter(p.Middlewares, p.Routes)}
}
