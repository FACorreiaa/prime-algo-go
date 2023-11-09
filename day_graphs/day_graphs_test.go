package day_graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func returnWeightedAdjacencyList() WeightedAdjacencyList {
	return WeightedAdjacencyList{
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
}
func TestBFSGraphList(t *testing.T) {
	//Graphs
	//BFSGraphList
	//     >(1)<--->(4) ---->(5)
	//    /          |       /|
	// (0)     ------|------- |
	//    \   v      v        v
	//     >(2) --> (3) <----(6)

	graphList := returnWeightedAdjacencyList()

	output, _ := BFSGraphList(graphList, 0, 6)
	expected := []int{0, 1, 4, 5, 6}

	assert.Equal(t, expected, output)
}

func TestBFSGraphMatrix(t *testing.T) {
	////matrix
	matrix := WeightedAdjacencyMatrix{
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

	output := BFSMatrixList(matrix, 0, 6)
	expected := []int{0, 1, 4, 5, 6}
	assert.Equal(t, expected, output)
}

func TestDijkstraList(t *testing.T) {
	//dijkstra
	//      (1) --- (4) ---- (5)
	//    /  |       |       /|
	// (0)   | ------|------- |
	//    \  |/      |        |
	//      (2) --- (3) ---- (6)
	testDijkstra := returnWeightedAdjacencyList()
	output := DijkstraList(0, 6, testDijkstra)
	expected := []int{0, 1, 4, 5, 6}
	assert.Equal(t, expected, output)
}
