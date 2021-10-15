package api

import (
	"MTS/auth/httpserver/api/auth/basicauth"
	"MTS/auth/httpserver/middlewares"
	"MTS/auth/httpserver/profiler"

	"github.com/go-chi/chi"
)

// App is a struct for application with all builtins
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

// ApplyEndpoints applies handlers for router paths
func (a *App) ApplyEndpoints(router *chi.Mux) {
	router.Use(middlewares.Logging)

	// login/logout endpoints
	router.Group(func(r chi.Router) {
		r.Get("/login", a.login)
		r.Get("/logout", a.logout)
	})

	router.Group(func(r chi.Router) {
		r.Use(middlewares.Debug)
		r.Mount("/debug", profiler.Profiler())
	})

	// i/me endpoints
	router.Group(func(r chi.Router) {
		r.Use(middlewares.UpdateCookies)
		r.Use(middlewares.HandleJWT)
		r.Get("/i", a.identificate)
		r.Get("/me", a.identificate)
	})
}
