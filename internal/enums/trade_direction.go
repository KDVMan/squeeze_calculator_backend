package enums

import "github.com/go-playground/validator/v10"

type TradeDirection string

const (
	TradeDirectionLong  TradeDirection = "long"
	TradeDirectionShort TradeDirection = "short"
)

func (object TradeDirection) String() string {
	return string(object)
}

func TradeDirectionValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(TradeDirection); ok {
		return enum.TradeDirectionValid()
	}

	return false
}

func (object TradeDirection) TradeDirectionValid() bool {
	switch object {
	case TradeDirectionLong, TradeDirectionShort:
		return true
	default:
		return false
	}
}
