package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/pkg/errors"
)

func Initialize(v *validator.Validate) error {
	if err := v.RegisterValidation("enum", enum); err != nil {
		return errors.Wrap(err, "not valid")
	}

	return nil
}

func enum(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(constant.Enum)
	return ok && value.Valid()
}
