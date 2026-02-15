package validation

import (
	"encoding/json"
	"errors"

	"github.com/KevenAbraham/first-crud-go/src/configurations/error_mapping"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
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
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(
	validation_err error,
) *error_mapping.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return error_mapping.NewBadRequestError("Invalid filed type")
	}

	if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []error_mapping.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := error_mapping.Causes {
				Message: e.Translate(transl),
				Field: e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return error_mapping.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	}

	return error_mapping.NewBadRequestError("Error trying to convert fields")
}
