package handlers

import (
	"errors"
	"net/http"

	"github.com/andersondalmina/go-api-skeleton/models"
	"github.com/andersondalmina/go-api-skeleton/security"
)

type returnLoginJSON struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

// LoginHandler asdfasd
func LoginHandler(userRepository models.UserRepositoryInterface, jwtm *security.JWTManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var bodyVars struct {
			Email    string
			Password string
		}

		err := requestToJSONObject(r, &bodyVars)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		user, err := userRepository.GetUserByEmail(bodyVars.Email)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		err = security.CompareHashAndPassword(user.Password, bodyVars.Password)
		if err != nil {
			HandleHTTPError(w, http.StatusUnauthorized, errors.New("Senha inválida"))
			return
		}

		token, err := jwtm.GenerateToken(user.ID)
		if err != nil {
			HandleHTTPError(w, http.StatusInternalServerError, errors.New("Error while Signing Token"))
			return
		}

		HandleHTTPSuccess(w, returnLoginJSON{
			Token: token,
			User: models.User{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			},
		})
	}
}
