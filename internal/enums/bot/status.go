package enums_bot

import "github.com/go-playground/validator/v10"

type Status string

const (
	StatusNew    Status = "new"
	StatusStart  Status = "start"
	StatusStop   Status = "stop"
	StatusDelete Status = "delete"
)

func (object Status) String() string {
	return string(object)
}

func StatusValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(Status); ok {
		return enum.StatusValid()
	}

	return false
}

func (object Status) StatusValid() bool {
	switch object {
	case StatusNew, StatusStart, StatusStop, StatusDelete:
		return true
	default:
		return false
	}
}
