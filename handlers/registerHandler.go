package handlers

import (
	"fmt"
	"net/http"
)

// RegisterHandler asdfasd
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test!")
}
