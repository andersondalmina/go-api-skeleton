package handlers

import (
	"fmt"
	"net/http"

	"github.com/andersondalmina/go-api-skeleton/models"
	"github.com/gorilla/context"
)

type returnProfileJSON struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ProfileHandler return data of logged user
func ProfileHandler(userRepository models.UserRepositoryInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := context.Get(r, contextUserIDKey)

		user, err := userRepository.GetUserByID(fmt.Sprintf("%s", userID))
		if err != nil {
			HandleHTTPError(w, 1, err)
		}

		HandleHTTPSuccess(w, returnProfileJSON{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
}
