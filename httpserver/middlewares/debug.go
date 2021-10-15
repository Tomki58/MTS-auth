package middlewares

import (
	"MTS/auth/httpserver/config"
	"net/http"
)

func Debug(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.Cfg.Debug {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}
