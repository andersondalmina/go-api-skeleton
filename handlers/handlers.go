package handlers

import (
	"encoding/json"
	"net/http"
)

func requestToJSONObject(req *http.Request, jsonDoc interface{}) error {
	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)
	return decoder.Decode(jsonDoc)
}

//ErrorMessage error message struct
type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// HandleHTTPError formats and returns errors
func HandleHTTPError(w http.ResponseWriter, errno int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errno)
	json.NewEncoder(w).Encode(&ErrorMessage{
		Code:    errno,
		Message: err.Error(),
	})
}

// HandleHTTPSuccess formats and return with content
func HandleHTTPSuccess(w http.ResponseWriter, data interface{}, status ...int) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		HandleHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	if len(status) == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(status[0])
	}
}

// HandleHTTPSuccessNoContent formats and return with no content
func HandleHTTPSuccessNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
