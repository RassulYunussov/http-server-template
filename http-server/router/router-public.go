package router

import (
	"template/http-server/middleware"

	"go.uber.org/fx"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

type PublicRouterPamameters struct {
	fx.In
	Middlewares []middleware.Middleware `group:"public-middlewares"`
	Routes      []Route                 `group:"public-routes"`
}

type PublicRouter struct {
	fx.Out
	Router *muxtrace.Router `name:"public-router"`
}

func NewPublicRouter(p PublicRouterPamameters) PublicRouter {
	return PublicRouter{Router: createRouter(p.Middlewares, p.Routes)}
}
