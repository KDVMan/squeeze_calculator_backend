package models_calculator_formula_preset

type DeleteRequestModel struct {
	ID uint `validate:"required,gt=0"`
}
