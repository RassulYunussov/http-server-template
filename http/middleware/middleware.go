package middleware

import (
	"fmt"
	"net/http"

	"go.uber.org/fx"
)

type Middleware interface {
	Middleware(next http.Handler) http.Handler
	Order() int
}

// ByOrder implements sort.Interface for []Middleware based on
// the Order
type ByOrder []Middleware

func (a ByOrder) Len() int           { return len(a) }
func (a ByOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByOrder) Less(i, j int) bool { return a[i].Order() < a[j].Order() }

func AsMiddleware(f any, tag string) any {
	return fx.Annotate(
		f,
		fx.As(new(Middleware)),
		fx.ResultTags(fmt.Sprintf(`group:"%s"`, tag)),
	)
}
