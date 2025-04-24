package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ResponseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Processing the Validation Error for the received fields value
func ValidationError(validationErrors *validator.ValidationErrors) *ResponseBody {
	var errorMsg []string
	for _, err := range *validationErrors {
		switch err.ActualTag() {
		case "required":
			message := fmt.Sprintf("%s is a required field", err.Field())
			errorMsg = append(errorMsg, message)

		case "email":
			message := fmt.Sprintf("Invalid value for %s field", err.Field())
			errorMsg = append(errorMsg, message)
		}
	}

	return &ResponseBody{
		Status:  "Error",
		Message: strings.Join(errorMsg, ", "),
	}
}

// Encoding the ReponseBody Struct into json response
func WriteJsonResponse(w http.ResponseWriter, status int, body interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(body)
}
