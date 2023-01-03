package http

import (
	"net/http"
	"template/config"

	"go.uber.org/fx"
	"go.uber.org/zap"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

type PrivateServer struct {
	*http.Server
}

func NewPrivateHTTPServer(lc fx.Lifecycle, log *zap.Logger, config config.Configuration, router *muxtrace.Router) *PrivateServer {
	return &PrivateServer{create(lc, log, config.PrivateHttpServer.Port, config.PrivateHttpServer.ReadTimeout, config.PrivateHttpServer.WriteTimeout, config.PrivateHttpServer.RequestTimeout, router)}
}
