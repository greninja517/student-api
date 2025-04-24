package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/greninja517/student-api/internal/http/types"
	"github.com/greninja517/student-api/internal/http/utils/response"
)

func CreateStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// populating the struct from the incoming json body
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)

		// catching the empty body error
		if errors.Is(err, io.EOF) {
			// write the response
			response.WriteJsonResponse(w, http.StatusBadRequest, &response.ResponseBody{
				Status:  "Error",
				Message: "Empty Body",
			})
			return
		}

		// catching the general error
		if err != nil {
			response.WriteJsonResponse(w, http.StatusBadRequest, &response.ResponseBody{
				Status:  "Error",
				Message: error.Error(err),
			})
			return
		}

		//validation of incoming request
		validate := validator.New()
		if err := validate.Struct(student); err != nil {
			// assertion
			validationError := err.(validator.ValidationErrors)
			response.WriteJsonResponse(w, http.StatusBadRequest, response.ValidationError(&validationError))
			return
		}

		// Resource Creation in no Error in Encountered
		slog.Info("Student Created Successfully", slog.String("Status", "Success"))
		response.WriteJsonResponse(w, http.StatusCreated, &response.ResponseBody{
			Status:  "OK",
			Message: "Resource Successfully Created",
		})

	}
}
