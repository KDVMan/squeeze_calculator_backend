package models_bot

type UpdateRequestModel struct {
	CalculatorPresetID        uint `json:"calculatorPresetId" validate:"required,gt=0"`
	CalculatorFormulaPresetID uint `json:"calculatorFormulaPresetId" validate:"required,gt=0"`
}
