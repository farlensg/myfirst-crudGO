package validation

import (
	"encoding/json"
	"errors"

	"github.com/farlensg/myfirst-crudGO/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationErr validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationErr) {
		// Retorna apenas a primeira mensagem de erro
		for _, e := range validation_err.(validator.ValidationErrors) {
			return rest_err.NewBadRequestError(e.Translate(transl))
		}
		return rest_err.NewBadRequestError("Validation error")
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
