package api

import (
	"MTS/auth/httpserver/serializer"
	"errors"
	"net/http"
)

type ContextKey string

const ContextUserKey ContextKey = "user"

func (a *App) identificate(w http.ResponseWriter, r *http.Request) {
	if name, ok := r.Context().Value("user").(string); ok {
		response := serializer.SerializeResponseJSON(name)
		if _, err := w.Write(response); err != nil {
			return
		}
	} else {
		err := errors.New("cannot fetch username from context")
		response := serializer.SerializeResponseJSON(err)
		http.Error(w, string(response), http.StatusInternalServerError)
		return
	}
}
