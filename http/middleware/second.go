package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

type SecondMiddleware struct {
	log *zap.Logger
}

func (*SecondMiddleware) Order() int {
	return 1
}

func (m *SecondMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.log.Info("touched SecondMiddleware")
		next.ServeHTTP(w, r)
	})
}

func NewSecondMiddleware(log *zap.Logger) *SecondMiddleware {
	return &SecondMiddleware{log}
}
