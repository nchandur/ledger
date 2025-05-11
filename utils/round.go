package utils

import "math"

func RoundToCent(val float64) float64 {
	val = math.Round(val * 100)

	return val / 100
}
