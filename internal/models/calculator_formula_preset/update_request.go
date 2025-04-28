package models_calculator_formula_preset

type UpdateRequestModel struct {
	ID       uint   `validate:"required,gt=0"`
	Name     string `validate:"required"`
	Filters  []FilterModel
	Formulas []FormulaModel
}
