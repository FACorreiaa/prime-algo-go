package day_search

func BSList(hstack []int64, needle int64) bool {
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
