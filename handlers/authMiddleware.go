package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/andersondalmina/go-api-skeleton/security"
	"github.com/gorilla/context"
)

const contextTokenKey = "tokenKey"
const contextUserIDKey = "userIDKey"

// AuthenticateMiddleware valida o token e filtra usuários não logados corretamente
func AuthenticateMiddleware(jwtm *security.JWTManager) func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		token := r.Header.Get("Authorization")
		if token == "" {
			HandleHTTPError(w, http.StatusUnauthorized, errors.New("Error no token is provided"))
			return
		}

		token = strings.Split(token, " ")[1]

		t, err := jwtm.ValidateToken(token)
		if err != nil {
			HandleHTTPError(w, http.StatusUnauthorized, err)
			return
		}

		context.Set(r, contextTokenKey, t.Token)
		context.Set(r, contextUserIDKey, t.UserID)
		next(w, r)
		context.Clear(r)
	}
}
