package models_calculator_preset

type DeleteRequestModel struct {
	ID uint `validate:"required,gt=0"`
}
