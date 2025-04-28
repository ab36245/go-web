package web

import (
	"net/http"
)

func New() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

type Server struct {
	mux *http.ServeMux
}

func (s *Server) Handler(pattern string) *Handler {
	handler := &Handler{}
	s.mux.Handle(pattern, handler)
	return handler
}

func (s *Server) Serve(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}

func (s *Server) ServeTLS(addr, certFile, keyFile string) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, s.mux)
}
