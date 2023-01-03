package handler

import (
	"io"
	"net/http"

	"go.uber.org/zap"
)

type EchoHandler struct {
	log *zap.Logger
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}

func (*EchoHandler) Methods() []string {
	return []string{http.MethodPost}
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.log.Error("Failed to handle request", zap.Error(err))
	}
}

func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{log}
}
