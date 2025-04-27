package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/greninja517/student-api/internal/http/types"
	"github.com/greninja517/student-api/internal/http/utils/response"
	"github.com/greninja517/student-api/internal/storage"
)

func NewStudent(storage storage.Storage) http.HandlerFunc {
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

		// inserting the received values in the database
		id, err := storage.CreateStudent(student.Name, student.Email)
		if err != nil {
			response.WriteJsonResponse(w, http.StatusInternalServerError, &response.ResponseBody{
				Status:  "Error",
				Message: "Failed to Write",
			})
			slog.Info("", slog.String("Error: ", err.Error()))
			return
		}

		// Resource Creation if no Error in Encountered
		slog.Info("Student Created Successfully", slog.String("ID", strconv.FormatInt(id, 10)), slog.String("Status", "Success"))
		response.WriteJsonResponse(w, http.StatusCreated, &response.ResponseBody{
			Status:  "OK",
			Message: "Resource Successfully Created",
		})

	}
}

func GetStudentById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		// parsing id to int64 type
		parsedID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Info("", slog.String("Error", err.Error()))
			response.WriteJsonResponse(w, http.StatusBadRequest, &response.ResponseBody{
				Status:  "Error",
				Message: "Invalid id value",
			})
			return
		}

		// getting the student info
		slog.Info("Retreiving the Student Info...", slog.String("ID", id))
		var student types.Student
		student, err = storage.GetStudent(parsedID)
		if err != nil {
			slog.Info("", "Error", err.Error())
			response.WriteJsonResponse(w, http.StatusInternalServerError, &response.ResponseBody{
				Status:  "Error",
				Message: "Not Found",
			})
			return
		}

		slog.Info("Retrieval Finished...", slog.String("Status", "Success"))
		response.WriteJsonResponse(w, http.StatusOK, student)

	}
}

func GetAllStudents(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Retriving all students...")

		// retreiving the students
		var students []types.Student
		students, err := storage.GetAll()
		if err != nil {
			slog.Info("", slog.String("Error", err.Error()))
			response.WriteJsonResponse(w, http.StatusInternalServerError, &response.ResponseBody{
				Status:  "Error",
				Message: "Error fetching all students",
			})
			return
		}

		slog.Info("Retrieval Finished...", slog.String("Status", "Success"))
		response.WriteJsonResponse(w, http.StatusOK, students)

	}
}
