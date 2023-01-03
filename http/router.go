package http

import (
	"net/http"
	"sort"

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

// AsRoute annotates the given constructor to state that
// it provides a route to the "routes" group.
func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

type Middleware interface {
	Middleware(next http.Handler) http.Handler
	Order() int
}

func AsMiddleware(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Middleware)),
		fx.ResultTags(`group:"middlewares"`),
	)
}

// ByOrder implements sort.Interface for []Middleware based on
// the Order
type ByOrder []Middleware

func (a ByOrder) Len() int           { return len(a) }
func (a ByOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByOrder) Less(i, j int) bool { return a[i].Order() < a[j].Order() }

func NewRouter(middlewares []Middleware, routes []Route) *muxtrace.Router {
	router := muxtrace.NewRouter()
	sort.Sort(ByOrder(middlewares))
	for _, m := range middlewares {
		router.Use(m.Middleware)
	}
	for _, r := range routes {
		router.Handle(r.Pattern(), r).Methods(r.Methods()...)
	}
	return router
}
