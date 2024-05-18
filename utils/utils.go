package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON[P any](r *http.Request, payload P) error {
	if r.Body == http.NoBody {
		return fmt.Errorf("please provide body in request")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON[P any](w http.ResponseWriter, status int, v P) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error_msg": err.Error()})
}
