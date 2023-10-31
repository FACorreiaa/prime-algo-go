package day_search

import (
	"testing"
)

func TestBS_List(t *testing.T) {
	tests := []struct {
		hstack   []int64
		needle   int64
		expected bool
	}{
		{[]int64{1, 2, 3, 4, 5}, 3, true},  // Element exists
		{[]int64{1, 2, 3, 4, 5}, 6, false}, // Element doesn't exist
		{[]int64{1, 2, 3, 4, 5}, 1, true},  // First element
		{[]int64{1, 2, 3, 4, 5}, 5, true},  // Last element
		{[]int64{1, 2, 3, 4, 5}, 2, true},  // Middle element
		{[]int64{}, 1, false},              // Empty slice
		{[]int64{1}, 1, true},              // Single element, found
		{[]int64{1}, 2, false},             // Single element, not found
	}

	for _, test := range tests {
		result := BSList(test.hstack, test.needle)
		if result != test.expected {
			t.Errorf("BS_List(%v, %v) = %v; want %v", test.hstack, test.needle, result, test.expected)
		}
	}
}

func TestTwoCrystallBalls(t *testing.T) {
	idx := 123 // Replace with your desired random index

	// Generate data with 10,000 elements, starting from idx to the end as true
	data := make([]bool, 10000)
	for i := idx; i < len(data); i++ {
		data[i] = true
	}

	result := TwoCrystallBalls(data)

	if result != idx {
		t.Errorf("Expected result to be %d, but got %d", idx, result)
	}

	// Test with a different array
	data2 := make([]bool, 821)
	result2 := TwoCrystallBalls(data2)

	if result2 != -1 {
		t.Errorf("Expected result2 to be -1, but got %d", result2)
	}
}

func BenchmarkBS_List(b *testing.B) {
	hstack := []int64{1, 2, 3, 4, 5}
	needle := int64(3)
	for i := 0; i < b.N; i++ {
		_ = BSList(hstack, needle)
	}
}
