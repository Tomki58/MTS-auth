package api

import (
	"MTS/auth/httpserver/api/auth/basicauth"
	"MTS/auth/httpserver/middlewares"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type App struct {
	Authenticator basicauth.BasicAuthorizator
}

// New creates App object with credentials in creds
func New(creds string) (*App, error) {
	auth, err := basicauth.New(creds)
	if err != nil {
		return nil, err
	}

	return &App{
		Authenticator: *auth,
	}, nil
}

func (a *App) ApplyEndpoints(router *chi.Mux) {
	router.Use(middleware.DefaultLogger)

	// login/logout endpoints
	router.Group(func(r chi.Router) {
		r.Get("/login", a.login)
		r.Get("/logout", a.logout)
	})

	// i/me endpoints
	router.Group(func(r chi.Router) {
		r.Use(middlewares.UpdateCookies)
		r.Use(middlewares.HandleJWT)
		r.Get("/i", a.identificate)
		r.Get("/me", a.identificate)
	})
}
