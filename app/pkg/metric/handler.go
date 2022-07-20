package metric

import (
	"net/http"
)

const URL = "/app/heartbeat"

type Handler struct {
}

type HandlerFunc interface {
	HandlerFunc(method, path string, handler http.HandlerFunc)
}

func (h *Handler) Register(router HandlerFunc) {
	router.HandlerFunc(http.MethodGet, URL, h.Heartbeat)
}

// Heartbeat
// @Summary Heartbeat metric
// @Tags Metrics
// @Success 204
// @Failure 400
// @Router /api/heartbeat [get]
func (h *Handler) Heartbeat(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(204)
}
