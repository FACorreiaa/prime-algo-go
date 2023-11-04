package day_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// Test case 1: Test with an unsorted array
	input := []int{4, 2, 7, 1, 9}
	expected := []int{1, 2, 4, 7, 9}
	output := BubbleSort(input)
	assert.Equal(t, expected, output)

	// Test case 2: Test with an already sorted array
	input = []int{1, 2, 3, 4, 5}
	expected = []int{1, 2, 3, 4, 5}
	output = BubbleSort(input)
	assert.Equal(t, expected, output)

	// Add more test cases as needed
}

func TestQueue(t *testing.T) {
	queue := NewQueue[int]()

	// Test Enqueue and Dequeue
	queue.Enqueue(1)
	queue.Enqueue(2)

	item, err := queue.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 1, item)

	item, err = queue.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 2, item)

	// Test Dequeue on an empty queue
	_, err = queue.Dequeue()
	assert.Error(t, err)
	assert.Equal(t, "queue is empty", err.Error())
}

func TestStack(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	item, err := stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 2, item)
}

func BenchmarkBubbleSort(b *testing.B) {
	input := []int{5, 3, 1, 4, 2}
	for i := 0; i < b.N; i++ {
		BubbleSort(input)
	}
}
