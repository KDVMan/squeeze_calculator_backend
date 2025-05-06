package enums_bot

import "github.com/go-playground/validator/v10"

type Action string

const (
	ActionStartAll  Action = "startAll"
	ActionStopAll   Action = "stopAll"
	ActionDeleteAll Action = "deleteAll"
)

func (object Action) String() string {
	return string(object)
}

func ActionValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(Action); ok {
		return enum.ActionValid()
	}

	return false
}

func (object Action) ActionValid() bool {
	switch object {
	case ActionStartAll, ActionStopAll, ActionDeleteAll:
		return true
	default:
		return false
	}
}
