package day_search

import "math"

func BS_List(hstack []int64, needle int64) bool {
	lo := 0
	hi := len(hstack)

	for lo < hi {
		m := (hi + lo) / 2
		v := hstack[m]
		switch {
		case v == needle:
			return true
		case v > needle:
			hi = m
		default:
			lo = m - 1
		}
	}
	return false
}

func Two_Crystall_Balls(breaks []bool) int {
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
