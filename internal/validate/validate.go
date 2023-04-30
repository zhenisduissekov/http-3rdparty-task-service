package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

var validate = validator.New()

type ErrorResponse struct {
	FailedField   string
	Tag           string
	Value         string
	ReceivedValue interface{}
}

func Validate(items interface{}) (errors []*ErrorResponse) {
	err := validate.Struct(items)
	if err != nil {
		log.Err(err).Msgf("could not validate items %+v", items)
		for _, err := range err.(validator.ValidationErrors) {
			if err != nil {
				log.Err(err).Msg("could not validate items")
			}
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			element.ReceivedValue = err.Value()
			errors = append(errors, &element)
		}
	}

	return errors
}
