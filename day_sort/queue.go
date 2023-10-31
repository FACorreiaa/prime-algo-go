package day_sort

import "reflect"

type NodeQueue[T any] struct {
	value T
	next  *NodeQueue[T]
}

type Queue[T any] struct {
	length int
	head   *NodeQueue[T]
	tail   *NodeQueue[T]
}

func newNode[T any](item T) *NodeQueue[T] {
	return &NodeQueue[T]{value: item}
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{length: 0, head: nil, tail: nil}
}

func (q *Queue[T]) enqueue(item T) {
	node := newNode(item)
	q.length++
	if q.tail == nil {
		q.head, q.tail = node, node
	}
	q.tail.next = node
	q.tail = node
}

func (q *Queue[T]) dequeue() (T, bool) {
	//fix later
	zeroValue := reflect.Zero(reflect.TypeOf((*T)(nil)).Elem()).Interface()

	if q.head == nil {
		return zeroValue, false
	}

	q.length--

	head := q.head
	q.head = q.head.next

	//free
	head.next = nil
	return head.value, true
}
func (q *Queue[T]) peek() (T, bool) {
	var zero T
	if q.head == nil {
		return zero, false
	}
	return q.head.value, true
}
