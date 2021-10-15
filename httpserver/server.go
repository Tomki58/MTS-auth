package httpserver

import (
	"MTS/auth/httpserver/api"
	"MTS/auth/httpserver/config"
	"net/http"

	"github.com/go-chi/chi"
)

// A Server is a wrapping structure over http.Server
// it contains a boolean var debug for enabling/disabling /debug endpoint
type Server struct {
	httpServer *http.Server
}

// New creates new Server and returns it
func New() (*Server, error) {

	subApp, err := api.New("./data/credentials.txt")
	if err != nil {
		return nil, err
	}
	router := chi.NewRouter()

	subApp.ApplyEndpoints(router)

	httpServer := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	return &Server{
		httpServer: &httpServer,
	}, nil
}

// ListenAndServe invokes Server.httpServer common ListenAndServe method
func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}

// Switch changes the state of debugging
func (s *Server) Switch() {
	config.Cfg.Switch()
}
