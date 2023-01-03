package router

import (
	"fmt"
	"net/http"
	"sort"
	"template/http/middleware"

	"go.uber.org/fx"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

// Route is an http.Handler that knows the mux pattern
// under which it will be registered.
type Route interface {
	http.Handler

	// Pattern reports the path at which this is registered.
	Pattern() string
	Methods() []string
}

func AsRoute(f any, tag string) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(fmt.Sprintf(`group:"%s"`, tag)),
	)
}

func NewRouter(middlewares []middleware.Middleware, routes []Route) *muxtrace.Router {
	router := muxtrace.NewRouter()
	sort.Sort(middleware.ByOrder(middlewares))
	for _, m := range middlewares {
		router.Use(m.Middleware)
	}
	for _, r := range routes {
		router.Handle(r.Pattern(), r).Methods(r.Methods()...)
	}
	return router
}
