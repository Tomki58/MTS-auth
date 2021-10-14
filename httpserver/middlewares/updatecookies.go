package middlewares

import (
	"MTS/auth/httpserver/cookies"
	"MTS/auth/httpserver/serializer"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func UpdateCookies(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fetching refresh cookie from the request
		refreshCookie, err := r.Cookie("refresh")
		if err != nil {
			err := errors.New("You are not logged in!")
			http.Error(w, string(serializer.SerializeResponseJSON(err)), http.StatusForbidden)
			return
		}

		// fetcing refresh cookie value as a tokenString
		tokenString := refreshCookie.Value
		token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte("MyBaseKey"), nil
		})
		if err != nil {
			http.Error(w, string(serializer.SerializeResponseJSON(err)), http.StatusInternalServerError)
			return
		}

		// creating fresh cookies
		freshCookies, err := cookies.CreateCookies(*token)
		if err != nil {
			http.Error(w, string(serializer.SerializeResponseJSON(err)), http.StatusInternalServerError)
			return
		}

		// updating cookies on the client's side and in the request
		for _, cookie := range freshCookies {
			r.AddCookie(cookie)
			http.SetCookie(w, cookie)
		}

		next.ServeHTTP(w, r)
	})
}
