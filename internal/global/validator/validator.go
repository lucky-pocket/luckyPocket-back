package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func Initialize() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("enum", enum)
		if err != nil {
			return
		}
	}
}

func enum(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(constant.Enum)
	return ok && value.Valid()
}
