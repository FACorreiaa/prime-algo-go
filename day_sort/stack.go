package day_sort

import "fmt"

type StackNode[T any] struct {
	value T
	prev  *StackNode[T]
}

type Stack[T any] struct {
	length int
	head   *StackNode[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{length: 0, head: nil}
}

func (s *Stack[T]) Push(item T) {
	node := &StackNode[T]{value: item}
	node.prev = s.head
	s.head = node
	s.length++
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T

	if s.length == 0 {
		return zero, fmt.Errorf("stack is empty")
	}

	head := s.head
	s.head = s.head.prev
	s.length--

	return head.value, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s.head == nil {
		return zero, fmt.Errorf("stack is empty")
	}

	return s.head.value, nil
}
