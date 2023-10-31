package day_search

import "math"

func TwoCrystallBalls(breaks []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jumpAmount

	for i < len(breaks) {
		if breaks[i] {
			break
		}
		i += jumpAmount
	}

	i -= jumpAmount

	for j := 0; j < jumpAmount && i < len(breaks); j++ {
		if breaks[i] {
			return i
		}
	}

	return -1
}
