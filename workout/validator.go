package workout

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func checkTitleLength(title string) error {
	validate := validator.New()
	if validate.Var(title, "required,max=75") != nil {
		return errors.New("Title must be less than 76 characters!")
	}
	return nil
}
