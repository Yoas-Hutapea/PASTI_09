package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

type ValidationMiddleware struct {
	validator *validator.Validate
}

func NewValidationMiddleware() *ValidationMiddleware {
	return &ValidationMiddleware{
		validator: validator.New(),
	}
}

func (vm *ValidationMiddleware) ValidateInput(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body into a map
		var requestBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Validate the request body
		err = vm.validator.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			var validationMessages []string
			for _, fieldError := range validationErrors {
				validationMessages = append(validationMessages, fieldError.Error())
			}
			response := map[string]interface{}{
				"errors": validationMessages,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Proceed to the next handler if validation passes
		next.ServeHTTP(w, r)
	})
}
