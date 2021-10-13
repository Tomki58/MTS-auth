package api

import (
	"MTS/auth/httpserver/cookies"
	"MTS/auth/httpserver/serializer"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func (a *App) login(w http.ResponseWriter, r *http.Request) {
	// Validate data in http.Request
	username, password, ok := r.BasicAuth()
	if !ok {
		response := serializer.SerializeResponseJSON(errors.New("incorrect authorization header"))
		http.Error(w, string(response), http.StatusForbidden)
		return
	}

	// check credentials
	if ok := a.Authenticator.CheckCredentials(username, password); !ok {
		response := serializer.SerializeResponseJSON(errors.New("invalid credentials"))
		http.Error(w, string(response), http.StatusForbidden)
		return
	}

	// create new JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Username: username,
	})

	// generate cookies
	cookies, err := cookies.CreateCookies(*token)
	if err != nil {
		response := serializer.SerializeResponseJSON(err)
		http.Error(w, string(response), http.StatusInternalServerError)
		return
	}

	// set cookies
	http.SetCookie(w, cookies[0])
	http.SetCookie(w, cookies[1])
	w.WriteHeader(http.StatusAccepted)
}

func (a *App) logout(w http.ResponseWriter, r *http.Request) {
	access := http.Cookie{Name: "access"}
	refresh := http.Cookie{Name: "refresh"}

	http.SetCookie(w, &access)
	http.SetCookie(w, &refresh)
}
