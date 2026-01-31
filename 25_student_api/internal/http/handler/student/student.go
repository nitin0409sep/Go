package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nitin0409sep/students-api/internal/storage"
	types "github.com/nitin0409sep/students-api/internal/types"
	"github.com/nitin0409sep/students-api/internal/utils/response"
)

// Create
func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)  // Parse request body JSON into a Go struct

		if errors.Is(err, io.EOF) { // If body is empty
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body"))) // Giving Custom Error Message we uses fmt.Errorf 
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// Request Validation
		if err := validator.New().Struct(student); err != nil {

			validateErrs := err.(validator.ValidationErrors)  // Type caste the errors for validateErrs function

			response.WriteJson(w, http.StatusBadRequest, response.ValidationErrors(validateErrs))
			return 
		}

		slog.Info("student info - ",student)
		
		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)


		slog.Info("user created successfully", slog.String("userId", fmt.Sprint(lastId)))
		
		if(err != nil) {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}
		
		response.WriteJson(w, http.StatusInternalServerError, lastId)
	}

}
