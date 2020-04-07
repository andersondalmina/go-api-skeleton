package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/andersondalmina/go-api-skeleton/models"
	"github.com/andersondalmina/go-api-skeleton/security"
)

type returnRegisterJSON struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// RegisterHandler asdfasd
func RegisterHandler(userRepository models.UserRepositoryInterface, jwtm *security.JWTManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var bodyVars struct {
			Name     string
			Email    string
			Password string
		}

		err := requestToJSONObject(r, &bodyVars)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		password, err := security.GenerateHash(bodyVars.Password)
		if err != nil {
			log.Fatal(err)
		}

		u := models.User{
			Name:     bodyVars.Name,
			Email:    bodyVars.Email,
			Password: password,
		}

		user, err := userRepository.CreateUser(u)
		if err != nil {
			log.Fatal(err)
		}

		token, err := jwtm.GenerateToken(user.ID)
		if err != nil {
			HandleHTTPError(w, http.StatusInternalServerError, errors.New("Error while Signing Token"))
			return
		}

		HandleHTTPSuccess(w, returnRegisterJSON{
			Token: token,
			User: models.User{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			},
		})
	}
}
