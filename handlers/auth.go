package handlers

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func SignToken(pass string) (string, error) {

	secret := []byte("peacock")

	jwtToken := jwt.New(jwt.SigningMethodHS256)

	signedToken, err := jwtToken.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func Auth(next http.HandlerFunc, signedToken string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var jwtFromCookie string

		cookie, err := r.Cookie("token")
		if err == nil {
			jwtFromCookie = cookie.Value
		}

		if signedToken != jwtFromCookie {
			http.Error(w, "Authentification required", http.StatusUnauthorized)
			return
		}
		next(w, r)
	})
}
