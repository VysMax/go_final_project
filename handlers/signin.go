package handlers

import (
	"VysMax/models"
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var (
		pass   models.Auth
		result models.AuthResponse
		buf    bytes.Buffer
	)

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err = json.Unmarshal(buf.Bytes(), &pass); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	switch pass.Password == os.Getenv("TODO_PASSWORD") {
	case true:
		secret := []byte("peacock")
		jwtToken := jwt.New(jwt.SigningMethodHS256)
		if err != nil {
			http.Error(w, "Failed to parse token", http.StatusUnauthorized)
		}
		signedToken, err := jwtToken.SignedString(secret)
		if err != nil {
			http.Error(w, "Failed to parse token", http.StatusUnauthorized)
		}
		result.Token = signedToken
	case false:
		result.Error = "Неверный пароль"
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
