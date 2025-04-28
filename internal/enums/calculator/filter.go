package enums_calculator

type Filter string

const (
	FilterLt    Filter = "lt"
	FilterLte   Filter = "lte"
	FilterGt    Filter = "gt"
	FilterGte   Filter = "gte"
	FilterEqual Filter = "equal"
)
