package services_helper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func GetCpu(minus int) int {
	threads := runtime.NumCPU() - minus

	if threads < 1 {
		threads = 1
	}

	return 2
	// return threads
}

func MustConvertStringToFloat64(value string) float64 {
	if result, err := strconv.ParseFloat(value, 64); err == nil {
		return result
	}

	return 0
}

func MustConvertByteToMd5(text []byte) string {
	hash := md5.Sum(text)

	return hex.EncodeToString(hash[:])
}

func MustConvertStringToMd5(text string) string {
	return MustConvertByteToMd5([]byte(text))
}

func Round(value float64, decimal int) float64 {
	if decimal == 0 {
		return math.Round(value)
	}

	multiplier := math.Pow(10, float64(decimal))

	return math.Round(value*multiplier) / multiplier
}

func Floor(value float64, decimal int) float64 {
	if decimal == 0 {
		return math.Floor(value)
	}

	multiplier := math.Pow(10, float64(decimal))

	return math.Floor(value*multiplier) / multiplier
}

func GetPercentFromMinMax(min float64, max float64, fix int) float64 {
	if min == 0 {
		return 0
	}

	result := ((max / min) * 100) - 100

	if fix > 0 {
		return Round(result, fix)
	}

	return result
}

func GetRandomFloatByInt(min float64, max float64, step float64) float64 {
	accuracy := CalculateAccuracy(step)
	minInt := int64(min * float64(accuracy))
	maxInt := int64(max * float64(accuracy))
	stepInt := int64(step * float64(accuracy))
	rangeInt := maxInt - minInt
	steps := rangeInt / stepInt
	randomStep := rand.Int63n(steps + 1)

	return float64(minInt+randomStep*stepInt) / float64(accuracy)
}

func CalculateAccuracy(step float64) int64 {
	stepStr := fmt.Sprintf("%g", step)
	decimalPlaces := 0

	if strings.Contains(stepStr, ".") {
		decimalPlaces = len(stepStr) - strings.Index(stepStr, ".") - 1
	}

	accuracy := int64(math.Pow10(decimalPlaces))

	if accuracy > 10 {
		return accuracy
	} else {
		return 10
	}
}

func GetRandomInt(min int64, max int64, step int64) int64 {
	if step <= 0 {
		panic("step must be positive")
	}

	if min > max {
		panic("min cannot be greater than max")
	}

	numSteps := (max-min)/step + 1
	randomStep := rand.Int63n(numSteps)

	return min + randomStep*step
}

func GetRangeFloatByInt(min float64, max float64, step float64) (int64, int64, int64, int64) {
	accuracy := CalculateAccuracy(step)
	minInt := int64(min * float64(accuracy))
	maxInt := int64(max * float64(accuracy))
	stepInt := int64(step * float64(accuracy))

	return minInt, maxInt, stepInt, accuracy
}

func GenerateRangeByStep(from int64, to int64, step int64) []int64 {
	var out []int64

	for value := from; value <= to; value += step {
		out = append(out, value)
	}

	return out
}

func MustConvertUnixMillisecondsToString(value int64) string {
	return time.UnixMilli(value).Format("02.01.2006 15:04:05.000")
}
