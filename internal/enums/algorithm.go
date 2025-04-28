package enums

import "github.com/go-playground/validator/v10"

type Algorithm string

const (
	AlgorithmRandom Algorithm = "random"
	AlgorithmGrid   Algorithm = "grid"
)

func AlgorithmValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(Algorithm); ok {
		return enum.AlgorithmValid()
	}

	return false
}

func (enum Algorithm) AlgorithmValid() bool {
	switch enum {
	case AlgorithmRandom, AlgorithmGrid:
		return true
	default:
		return false
	}
}
