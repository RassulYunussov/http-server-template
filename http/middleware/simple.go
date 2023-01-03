package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

type SimpleMiddleware struct {
	log *zap.Logger
}

func (*SimpleMiddleware) Order() int {
	return 0
}

func (m *SimpleMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.log.Info("touched SimpleMiddleware")
		next.ServeHTTP(w, r)
	})
}

func NewSimpleMiddleware(log *zap.Logger) *SimpleMiddleware {
	return &SimpleMiddleware{log}
}
