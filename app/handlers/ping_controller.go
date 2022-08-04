package handlers

import "net/http"

type PingHandler interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type pingHandler struct {
}

func NewPingHandler() PingHandler {
	return &pingHandler{}
}

func (s *pingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
