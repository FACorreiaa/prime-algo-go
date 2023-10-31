package day_sort

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

func (s *Stack[T]) Pop() (T, bool) {
	var zero T

	if s.length == 0 {
		return zero, false
	}

	head := s.head
	s.head = s.head.prev
	s.length--

	return head.value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if s.head == nil {
		return zero, false
	}

	return s.head.value, true
}
