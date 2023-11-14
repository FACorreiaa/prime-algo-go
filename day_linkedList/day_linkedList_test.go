package day_linkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {

	doubleLinkedList := NewLinkedList[int]()
	// Append operations
	doubleLinkedList.Append(5)
	doubleLinkedList.Append(7)
	doubleLinkedList.Append(9)

	second, _ := doubleLinkedList.Get(2)
	removeFirst, _ := doubleLinkedList.RemoveAt(1)
	println(second)
	// Assertions
	assert.Equal(t, 9, second)
	assert.Equal(t, 7, removeFirst)
	assert.Equal(t, 2, doubleLinkedList.Length)

	// Append and Remove operations
	doubleLinkedList.Append(11)
	//assert.Equal(t, 9, doubleLinkedList.RemoveAt(1))
	//assert.Equal(t, -1, doubleLinkedList.Remove(9))
	//assert.Equal(t, 5, doubleLinkedList.RemoveAt(0))
	//assert.Equal(t, 11, doubleLinkedList.RemoveAt(0))
	//assert.Equal(t, 0, doubleLinkedList.Length)

	// Prepend operations
	doubleLinkedList.Prepend(5)
	doubleLinkedList.Prepend(7)
	doubleLinkedList.Prepend(9)

	// Assertions
	//assert.Equal(t, 5, doubleLinkedList.Get(2))
	//assert.Equal(t, 9, doubleLinkedList.Get(0))
	//assert.Equal(t, 9, doubleLinkedList.Remove(9))
	//assert.Equal(t, 2, doubleLinkedList.Length)
	//assert.Equal(t, 7, doubleLinkedList.Get(0))

}
