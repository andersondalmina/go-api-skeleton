package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler asdfasdf
func HomeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome home!")
	}
}
