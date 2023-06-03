package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ValidationMiddleware struct {
	validator *validator.Validate
}

func NewValidationMiddleware() *ValidationMiddleware {
	// Initialize and configure the ValidationMiddleware instance
	return &ValidationMiddleware{
		// Initialize fields or dependencies
	}
}

func (vm *ValidationMiddleware) ValidateUserInput(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Use the validator to validate the request form values
		err = vm.validator.Struct(r.Form)
		if err != nil {
			// If there are validation errors, retrieve and handle them
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				// Map validation errors to a more readable format
				errorMap := make(map[string]string)
				for _, e := range validationErrors {
					errorMap[e.Field()] = e.Tag()
				}

				// Convert the errorMap to JSON
				jsonErrors, err := json.Marshal(errorMap)
				if err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}

				// Return the validation errors as a JSON response
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write(jsonErrors)
				return
			}
		}

		// Call the next handler if validation passes
		next(w, r)
	}
}