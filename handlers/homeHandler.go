package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler asdfasdf
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
