package app

import (
	"math"

	"github.com/a-h/round"
)

// EstimateIterations -
func EstimateIterations(substring string, index int) int {
	// single iteration needed for empty match
	searchSpace := 1

	digitsRemaining := len(substring)
	if digitsRemaining == 0 {
		return searchSpace
	}

	// no iterations needed for "xrb_"
	if index < 4 {
		digitsRemaining = int(math.Max(float64(digitsRemaining-(4-index)), 0))
		if digitsRemaining == 0 {
			return searchSpace
		}
	}

	// 2 possibilities for the first address digit (1 or 3)
	if index <= 4 {
		searchSpace = searchSpace * 2
		digitsRemaining = digitsRemaining - 1
		if digitsRemaining == 0 {
			return searchSpace
		}
	}

	// charset of 32 for the rest of the digits
	searchSpace = searchSpace * int(round.AwayFromZero(math.Pow(32, float64(digitsRemaining)), 0))

	// discovery time is assumed to be the mean
	return searchSpace / 2
}
