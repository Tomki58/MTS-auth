package httpserver

import (
	"MTS/auth/httpserver/api"
	"net/http"

	"github.com/go-chi/chi"
)

func New() (*http.Server, error) {

	subApp, err := api.New("./auth/credentials.txt")
	if err != nil {
		return nil, err
	}
	router := chi.NewRouter()
	subApp.ApplyEndpoints(router)

	httpServer := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	return &httpServer, nil
}
