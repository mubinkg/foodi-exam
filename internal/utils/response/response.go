package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOk    = "Ok"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errors validator.ValidationErrors) Response {
	var errMessages []string

	for _, err := range errors {
		switch err.ActualTag() {
		case "required":
			errMessages = append(errMessages, fmt.Sprintf("filed %s is required", err.Field()))
		default:
			errMessages = append(errMessages, fmt.Sprintf("filed %s is invalid", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMessages, ", "),
	}
}
