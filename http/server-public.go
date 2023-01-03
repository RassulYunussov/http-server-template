package http

import (
	"net/http"
	"template/config"

	"go.uber.org/fx"
	"go.uber.org/zap"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

type PublicServer struct {
	*http.Server
}

func NewPublicHTTPServer(lc fx.Lifecycle, log *zap.Logger, config config.Configuration, router *muxtrace.Router) *PublicServer {
	return &PublicServer{create(lc, log, config.PublicHttpServer.Port, config.PublicHttpServer.ReadTimeout, config.PublicHttpServer.WriteTimeout, config.PublicHttpServer.RequestTimeout, router)}
}
