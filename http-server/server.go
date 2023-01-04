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
)

func create(lc fx.Lifecycle, log *zap.Logger, httpServer config.HttpServer, handler http.Handler) *http.Server {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", httpServer.Port),
		ReadTimeout:  time.Second * time.Duration(httpServer.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(httpServer.WriteTimeout),
		Handler:      http.TimeoutHandler(handler, time.Second*time.Duration(httpServer.RequestTimeout), "Timeout!"),
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
						log.Info(fmt.Sprintf("HTTP server at %d closed", httpServer.Port))
					default:
						log.Error(fmt.Sprintf("HTTP server at %d error", httpServer.Port), zap.Error(err))
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

var HttpServersModule = fx.Module("http-servers",
	PublicServerModule,
	PrivateServerModule,
)
