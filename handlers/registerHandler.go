package handlers

import (
	"log"
	"net/http"

	"github.com/andersondalmina/go-api-skeleton/models"
	"github.com/andersondalmina/go-api-skeleton/security"
)

// RegisterHandler asdfasd
func RegisterHandler(userRepository models.UserRepositoryInterface) func(w http.ResponseWriter, r *http.Request) {
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

		u, err = userRepository.CreateUser(u)
		if err != nil {
			log.Fatal(err)
		}

		HandleHTTPSuccess(w, map[string]string{"id": u.ID})
	}
}
