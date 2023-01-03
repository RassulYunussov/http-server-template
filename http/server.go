package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func create(lc fx.Lifecycle, log *zap.Logger, port int16, readTimeout int, writeTimeout int, requestTimeout int, handler http.Handler) *http.Server {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  time.Second * time.Duration(readTimeout),
		WriteTimeout: time.Second * time.Duration(writeTimeout),
		Handler:      http.TimeoutHandler(handler, time.Second*time.Duration(requestTimeout), "Timeout!"),
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
						log.Info(fmt.Sprintf("HTTP server at %d closed", port))
					default:
						log.Error(fmt.Sprintf("HTTP server at %d error", port), zap.Error(err))
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
