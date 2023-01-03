package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"template/config"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

func NewHTTPServer(lc fx.Lifecycle, log *zap.Logger, config config.Configuration, router *muxtrace.Router) *http.Server {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Httpserver.Port),
		ReadTimeout:  time.Second * time.Duration(config.Httpserver.Readtimeout),
		WriteTimeout: time.Second * time.Duration(config.Httpserver.Writetimeout),
		Handler:      http.TimeoutHandler(router, time.Second*time.Duration(config.Httpserver.Requesttimeout), "Timeout!"),
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting HTTP server", zap.String("at", srv.Addr))
			go func() {
				if err := srv.Serve(ln); err != nil {
					switch err {
					case http.ErrServerClosed:
						log.Info("HTTP server closed")
					default:
						log.Error("HTTP server error", zap.Error(err))
					}
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
