package validation

import (
	"gopkg.in/go-playground/validator.v9"
	"github.com/Yoas-Hutapea/Microservice_09/auth/models"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type ValidationError map[string]string

func (v ValidationError) Error() string {
	var errMsg string
	for _, err := range v {
		errMsg += err + "\n"
	}
	return errMsg
}

type AuthValidator struct{}

func NewAuthValidator() *AuthValidator {
	return &AuthValidator{}
}

func (av *AuthValidator) ValidateLoginRequest(request *models.LoginRequest) error {
	err := validate.Struct(request)
	if err != nil {
		validationErrors := ValidationError{}
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "NIK":
				switch err.Tag() {
				case "required":
					validationErrors["nik"] = "nik harus diisi"
				case "min":
					validationErrors["nik"] = "nik harus 16 angka"
				case "max":
					validationErrors["nik"] = "nik harus 16 angka"
				}
			case "Password":
				switch err.Tag() {
				case "required":
					validationErrors["password"] = "Password harus diisi"
				case "min":
					validationErrors["password"] = "Password minimal 8 karakter"
				}
			}
		}
		return validationErrors
	}
	return nil
}
