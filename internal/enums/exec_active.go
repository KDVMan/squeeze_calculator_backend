package enums

import "github.com/go-playground/validator/v10"

type ExecActive string

const (
	ExecActiveBot       ExecActive = "bot"
	ExecActiveCalculate ExecActive = "calculate"
)

func ExecActiveValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(ExecActive); ok {
		return enum.ExecActiveValid()
	}

	return false
}

func (enum ExecActive) ExecActiveValid() bool {
	switch enum {
	case ExecActiveBot, ExecActiveCalculate:
		return true
	default:
		return false
	}
}
