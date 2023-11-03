package main

import (
	"fmt"
	leet_code_trainning "github.com/FACorreiaa/prime-algo-go/LeetCodeTrainning"
	"github.com/FACorreiaa/prime-algo-go/day_graphs"
	"github.com/FACorreiaa/prime-algo-go/day_heap"
	"github.com/FACorreiaa/prime-algo-go/day_linkedList"
	"github.com/FACorreiaa/prime-algo-go/day_maps"
	"github.com/FACorreiaa/prime-algo-go/day_quickSort"
	"github.com/FACorreiaa/prime-algo-go/day_recursion"
	"github.com/FACorreiaa/prime-algo-go/day_search"
	"github.com/FACorreiaa/prime-algo-go/day_sort"
	"github.com/FACorreiaa/prime-algo-go/day_tree"
	"github.com/FACorreiaa/prime-algo-go/day_tree_search"
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

func intSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func matrixIsEqual(a, b day_graphs.WeightedAdjacencyMatrix) bool {
	if len(a.Values) != len(b.Values) {
		return false
	}

	for i := range a.Values {
		if len(a.Values[i]) != len(b.Values[i]) {
			return false
		}

		for j := range a.Values[i] {
			if a.Values[i][j] != b.Values[i][j] {
				return false
			}
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

	// Test the Compare function BinaryTreeComparison
	tree1 := &day_tree.BinaryNode[int]{Value: 1}
	tree1.Left = &day_tree.BinaryNode[int]{Value: 2}
	tree1.Right = &day_tree.BinaryNode[int]{Value: 3}
	tree2 := &day_tree.BinaryNode[int]{Value: 1}
	tree2.Left = &day_tree.BinaryNode[int]{Value: 2}
	tree2.Right = &day_tree.BinaryNode[int]{Value: 3}

	areEqual := day_tree.Compare(tree1, tree2)
	fmt.Println("Are the trees equal?", areEqual) // Should print "Are the trees equal? true"

	// Test the BreadthFirstSearch function BFS
	head := &day_tree.BinaryNode[int]{Value: 1}
	head.Left = &day_tree.BinaryNode[int]{Value: 2}
	head.Right = &day_tree.BinaryNode[int]{Value: 3}
	head.Left.Left = &day_tree.BinaryNode[int]{Value: 4}
	head.Left.Right = &day_tree.BinaryNode[int]{Value: 5}

	needle := 3
	found := day_tree.BreadthFirstSearch(head, needle)
	fmt.Println("Found needle?", found) // Should print "Found needle? true"
	//DepthFirstSearchOnBST
	less := func(a, b int) bool {
		return a < b
	}

	foundDFS := day_tree.DFS(head, needle, less)
	fmt.Println("Found needle?", foundDFS)
	//tree search
	inOrder := day_tree_search.InOrderSearch(head)
	fmt.Println("In order search", inOrder)

	postOrder := day_tree_search.PostOrderSearch(head)
	fmt.Println("Post order search", postOrder)

	preOrder := day_tree_search.PreOrderSearch(head)
	fmt.Println("Post order search", preOrder)

	//Heap
	//Tries
	trie := day_heap.NewTrie()
	trie.Insert("apple")
	trie.Insert("app")
	trie.Insert("banana")

	fmt.Println("Before Deletion:")
	fmt.Println(trie.Root) // Trie structure

	trie.Delete("app")

	fmt.Println("After Deletion:")
	fmt.Println(trie.Root) // Trie structure

	heap := day_heap.NewMinHeap()
	heap.Insert(1)
	heap.Insert(2)
	heap.Insert(3)
	heap.Delete()
	fmt.Println(heap.Length) // Heap structure
	fmt.Println(heap.Data)   // Heap structure

	//Graphs
	//BFSGraphList
	//     >(1)<--->(4) ---->(5)
	//    /          |       /|
	// (0)     ------|------- |
	//    \   v      v        v
	//     >(2) --> (3) <----(6)

	graphList := day_graphs.WeightedAdjacencyList{
		{
			{To: 1, Weight: 3},
			{To: 2, Weight: 1},
		},
		{
			{To: 4, Weight: 1},
		},
		{
			{To: 3, Weight: 7},
		},
		{},
		{
			{To: 1, Weight: 1},
			{To: 3, Weight: 5},
			{To: 5, Weight: 2},
		},
		{
			{To: 2, Weight: 18},
			{To: 6, Weight: 1},
		},
		{
			{To: 3, Weight: 1},
		},
	}

	graphListPoints, ok := day_graphs.BFSGraphList(graphList, 6, 0)
	if !ok {
		fmt.Println("Its not on matrix")
	} else {
		fmt.Println(graphListPoints)
	}

	////matrix
	matrix := day_graphs.WeightedAdjacencyMatrix{
		// Define the values for the matrix here
		Values: [][]int{
			{0, 3, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 0},
			{0, 0, 7, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 5, 0, 2, 0},
			{0, 0, 18, 0, 0, 0, 1},
			{0, 0, 0, 1, 0, 0, 1},
		},
	}

	sampleMatrix := day_graphs.WeightedAdjacencyMatrix{
		// Define the values for the matrix here
		Values: [][]int{
			{0, 3, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 0},
			{0, 0, 7, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 5, 0, 2, 0},
			{0, 0, 18, 0, 0, 0, 1},
			{0, 0, 0, 1, 0, 0, 1},
		},
	}
	matrisIsEqual := matrixIsEqual(matrix, sampleMatrix)
	fmt.Println(matrisIsEqual)
	graphMatrix := day_graphs.BFSMatrixList(sampleMatrix, 6, 0)
	if graphMatrix == nil {
		fmt.Println("Its not on matrix")
	} else {
		fmt.Println("Its on matrix")

		fmt.Println(graphMatrix)
	}

	//dijkstra
	//      (1) --- (4) ---- (5)
	//    /  |       |       /|
	// (0)   | ------|------- |
	//    \  |/      |        |
	//      (2) --- (3) ---- (6)
	testDijkstra := day_graphs.WeightedAdjacencyList{
		{
			{To: 1, Weight: 3},
			{To: 2, Weight: 1},
		},
		{
			{To: 0, Weight: 3},
			{To: 2, Weight: 4},
			{To: 4, Weight: 1},
		},
		{
			{To: 1, Weight: 4},
			{To: 3, Weight: 7},
			{To: 0, Weight: 1},
		},
		{
			{To: 2, Weight: 7},
			{To: 4, Weight: 5},
			{To: 6, Weight: 1},
		},
		{
			{To: 1, Weight: 1},
			{To: 3, Weight: 5},
			{To: 5, Weight: 2},
		},
		{
			{To: 6, Weight: 1},
			{To: 4, Weight: 2},
			{To: 2, Weight: 18},
		},
		{
			{To: 3, Weight: 1},
			{To: 5, Weight: 1},
		},
	}

	sample := []int{0, 1, 4, 5, 6}
	listResult := day_graphs.DijkstraList(0, 6, testDijkstra)
	if intSlicesEqual(listResult, sample) {
		fmt.Println("listResult Result correct")
	} else {
		fmt.Println("listResult Result incorrect")
	}

	//Map
	day_maps.DisplayMap()

	//LRU

	//Prims
	primsList := day_graphs.WeightedAdjacencyList{
		{
			{To: 2, Weight: 1},
			{To: 1, Weight: 3},
		},
		{
			{To: 0, Weight: 3},
			{To: 4, Weight: 1},
			{To: 3, Weight: 3},
		},
		{
			{To: 0, Weight: 1},
			{To: 3, Weight: 7},
		},
		{
			{To: 6, Weight: 1},
			{To: 1, Weight: 3},
			{To: 2, Weight: 7},
		},
		{
			{To: 1, Weight: 1},
			{To: 5, Weight: 2},
		},
		{
			{To: 4, Weight: 2},
			{To: 6, Weight: 1},
		},
		{
			{To: 5, Weight: 1},
			{To: 3, Weight: 1},
		},
	}
	// Call your PrimsList function
	mst := leet_code_trainning.PrimsList(primsList)

	// Print the result
	for i, edges := range mst {
		fmt.Printf("Node %d:\n", i)
		for _, edge := range edges {
			fmt.Printf("  To: %d, Weight: %d\n", edge.To, edge.Weight)
		}
	}

	fmt.Println("PrimsList Test Complete")
}
