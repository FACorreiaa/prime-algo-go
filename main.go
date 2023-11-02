package main

import (
	"fmt"
	"github.com/FACorreiaa/prime-algo-go/day_linkedList"
	"github.com/FACorreiaa/prime-algo-go/day_quickSort"
	"github.com/FACorreiaa/prime-algo-go/day_recursion"
	"github.com/FACorreiaa/prime-algo-go/day_search"
	"github.com/FACorreiaa/prime-algo-go/day_sort"
	"math/rand"
	"time"
)

func equalSlices(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

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

	//Stack
	s := day_sort.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s)

	//Queue
	q := day_sort.NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q)

	//Maze Solver
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	mazeResult := []day_recursion.Point{
		{10, 0},
		{10, 1},
		{10, 2},
		{10, 3},
		{10, 4},
		{9, 4},
		{8, 4},
		{7, 4},
		{6, 4},
		{5, 4},
		{4, 4},
		{3, 4},
		{2, 4},
		{1, 4},
		{1, 5},
	}

	result := day_recursion.MazeSolver(maze, "x", day_recursion.Point{X: 10, Y: 10}, day_recursion.Point{X: 1, Y: 5})
	if equalSlices(day_recursion.DrawPath(maze, result), day_recursion.DrawPath(maze, mazeResult)) {
		fmt.Println("Slices are equal.")
	} else {
		fmt.Println("Slices are not equal.")
	}
	//Quick Sort
	day_quickSort.QuickSort(unSortedList)

	// Check if the list is sorted
	if isSorted(unSortedList) {
		fmt.Println("List is sorted")
	} else {
		fmt.Println("List is not sorted")
	}

	//double linked list
	doubleLinkedList := day_linkedList.NewLinkedList[int]()
	doubleLinkedList.Append(5)
	doubleLinkedList.Append(6)
	doubleLinkedList.Append(7)

	val, err := doubleLinkedList.Get(2)
	if err == nil && val == 9 {
		fmt.Println("Get(2) =>", val)
	}

	val, err = doubleLinkedList.RemoveAt(1)
	if err == nil && val == 7 {
		fmt.Println("RemoveAt(1) =>", val)
	}

	length := doubleLinkedList.Length
	fmt.Println("Length() =>", length)

	doubleLinkedList.Append(11)
	val, err = doubleLinkedList.RemoveAt(1)
	if err == nil && val == 9 {
		fmt.Println("RemoveAt(1) =>", val)
	}

	val, _ = doubleLinkedList.Remove(9)
	fmt.Println("Remove(9) =>", val)

	val, err = doubleLinkedList.RemoveAt(0)
	if err == nil && val == 5 {
		fmt.Println("RemoveAt(0) =>", val)
	}

	val, err = doubleLinkedList.RemoveAt(0)
	if err == nil && val == 11 {
		fmt.Println("RemoveAt(0) =>", val)
	}

	length = doubleLinkedList.Length
	fmt.Println("Length() =>", length)

	doubleLinkedList.Prepend(5)
	doubleLinkedList.Prepend(7)
	doubleLinkedList.Prepend(9)

	val, err = doubleLinkedList.Get(2)
	if err == nil && val == 5 {
		fmt.Println("Get(2) =>", val)
	}

	val, err = doubleLinkedList.Get(0)
	if err == nil && val == 9 {
		fmt.Println("Get(0) =>", val)
	}

	val, err = doubleLinkedList.Remove(9)
	fmt.Println("Remove(9) =>", val)

	length = doubleLinkedList.Length
	fmt.Println("Length() =>", length)

	val, err = doubleLinkedList.Get(0)
	if err == nil && val == 7 {
		fmt.Println("Get(0) =>", val)
	}

	//tree

}
