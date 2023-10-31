package main

import (
	"fmt"
	"github.com/FACorreiaa/prime-algo-go/day_search"
	"github.com/FACorreiaa/prime-algo-go/day_sort"
	"math/rand"
	"time"
)

func main() {
	//Binary Search
	println("Binary Search")
	values := []int64{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	bs := day_search.BSList(values, 99)
	println("Binary search value: ", bs)

	//Two Crystal Balls
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomInt := random.Intn(10000)
	randomFloat := random.Float64() * float64(randomInt)
	firstDataSource := make([]bool, 10000)
	secondDataSource := make([]bool, 821)
	for i := int(randomFloat); i < len(firstDataSource); i++ {
		firstDataSource[i] = true
	}
	firstExample := day_search.TwoCrystallBalls(firstDataSource)
	secondExample := day_search.TwoCrystallBalls(secondDataSource)
	if firstExample == 0 {
		fmt.Println("Pass: Result matches the random index.")
	} else {
		fmt.Println("Fail: Result does not match the random index.")
	}

	if secondExample != -1 {
		fmt.Printf("Expected result2 to be -1, but got %d\n", secondExample)
	}

	//Bubble Sort
	unSortedList := []int{69, 420, 201, 111, 980, 50, 60, 70}
	sortedList := day_sort.BubbleSort(unSortedList)
	fmt.Println(sortedList)

	//Queue
	s := day_sort.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s)

	//Stack

	//Maze Solver
}
