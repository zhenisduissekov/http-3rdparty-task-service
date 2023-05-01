package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

var validate = validator.New()

const (
	ItemsForToValidateMsg    = "items to validate: %+v"
	CouldNotValidateItemsMsg = "could not validate items"
)

type ErrorResponse struct {
	FailedField   string
	Tag           string
	Value         string
	ReceivedValue interface{}
}

func Validate(items interface{}) (errors []*ErrorResponse) {
	err := validate.Struct(items)
	if err != nil {
		log.Err(err).Msgf(ItemsForToValidateMsg, items)
		for _, err := range err.(validator.ValidationErrors) {
			if err != nil {
				log.Err(err).Msg(CouldNotValidateItemsMsg)
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
