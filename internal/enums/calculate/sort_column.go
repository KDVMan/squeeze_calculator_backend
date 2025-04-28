package enums_calculate

import "github.com/go-playground/validator/v10"

type SortColumn string

const (
	SortColumnBind                         SortColumn = "bind"
	SortColumnPercentIn                    SortColumn = "percentIn"
	SortColumnPercentOut                   SortColumn = "percentOut"
	SortColumnStopTime                     SortColumn = "stopTime"
	SortColumnStopPercent                  SortColumn = "stopPercent"
	SortColumnTotal                        SortColumn = "total"
	SortColumnTotalStops                   SortColumn = "totalStops"
	SortColumnTotalStopsMinus              SortColumn = "totalStopsMinus"
	SortColumnTotalStopsPlus               SortColumn = "totalStopsPlus"
	SortColumnTotalTakes                   SortColumn = "totalTakes"
	SortColumnTotalTakesPlus               SortColumn = "totalTakesPlus"
	SortColumnTotalProfitPercent           SortColumn = "totalProfitPercent"
	SortColumnTotalCumulativeProfitPercent SortColumn = "totalCumulativeProfitPercent"
	SortColumnMaxDrawdownPercent           SortColumn = "maxDrawdownPercent"
	SortColumnMaxTimeDeal                  SortColumn = "maxTimeDeal"
	SortColumnInOutRatio                   SortColumn = "inOutRatio"
	SortColumnCoefficient                  SortColumn = "coefficient"
	SortColumnCoefficientPlus              SortColumn = "coefficientPlus"
	SortColumnWinRate                      SortColumn = "winRate"
	SortColumnWinRatePlus                  SortColumn = "winRatePlus"
	SortColumnScore                        SortColumn = "score"
)

func (object SortColumn) String() string {
	return string(object)
}

func SortColumnValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(SortColumn); ok {
		return enum.SortColumnValid()
	}

	return false
}

func (object SortColumn) SortColumnValid() bool {
	switch object {
	case SortColumnPercentIn, SortColumnPercentOut, SortColumnBind, SortColumnStopTime, SortColumnStopPercent, SortColumnTotal, SortColumnTotalStops,
		SortColumnTotalStopsMinus, SortColumnTotalStopsPlus, SortColumnTotalTakes, SortColumnTotalTakesPlus, SortColumnTotalProfitPercent, SortColumnTotalCumulativeProfitPercent,
		SortColumnMaxDrawdownPercent, SortColumnMaxTimeDeal, SortColumnInOutRatio, SortColumnCoefficient, SortColumnCoefficientPlus, SortColumnWinRate, SortColumnWinRatePlus, SortColumnScore:
		return true
	default:
		return false
	}
}
