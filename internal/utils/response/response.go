package response

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `"json:"status" bson:"status"`
	Error  string `json:"error" bson:"error"`
}

func WriteJson(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func VlidateError(errs validator.ValidationErrors) Response {
	var errorMessages []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errorMessages = append(errorMessages, err.Field()+" is required")
		case "email":
			errorMessages = append(errorMessages, err.Field()+" must be a valid email")
		default:
			errorMessages = append(errorMessages, err.Field()+" is invalid")
		}
	}

	return Response{
		Status: "error",
		Error:  strings.Join(errorMessages, ", "),
	}
}
