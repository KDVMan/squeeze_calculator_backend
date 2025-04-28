package models_calculator_preset

type EditRequestModel struct {
	ID   uint   `validate:"required,gt=0"`
	Name string `validate:"required"`
}
