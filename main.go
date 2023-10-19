package main

import "FACorreiaa/prime-algo-course/day_search"

func main() {
	println("Binary Search")
	values := []int64{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	bs := day_search.BS_List(values, 99)
	println("Binary seaarch value: ", bs)
	println("Two Crystall Balls")

}
