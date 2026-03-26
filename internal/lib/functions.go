package lib

import (
	"math"
)

func MinutesToRoundedHours(minutes int) float64 {
	roundedMinutes := int(math.Round(float64(minutes)/15.0)) * 15
	hours := float64(roundedMinutes) / 60.0
	return math.Trunc(hours*100) / 100
}
