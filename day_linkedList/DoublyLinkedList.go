package day_linkedList

import (
	"fmt"
	"reflect"
)

type LinkedListNode[T any] struct {
	value T
	next  *LinkedListNode[T]
	prev  *LinkedListNode[T]
}

type DoubleLinkedList[T any] struct {
	Length int
	head   *LinkedListNode[T]
	tail   *LinkedListNode[T]
}

func newNode[T any](item T) *LinkedListNode[T] {
	return &LinkedListNode[T]{value: item}
}

func NewLinkedList[T any]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{Length: 0, head: nil, tail: nil}
}

func isEqual[T any](a, b T) bool {
	return reflect.DeepEqual(a, b)
}

func (d *DoubleLinkedList[T]) Prepend(item T) {
	node := newNode(item)
	d.Length++

	if d.head == nil {
		d.head = node
		return
	}

	node.next = d.head
	d.head.prev = node
	d.head = node
}

func (d *DoubleLinkedList[T]) Append(item T) {
	d.Length++
	node := newNode(item)
	if d.tail == nil {
		d.head, d.tail = node, node
		return
	}
	node.prev = d.tail
	d.tail.next = node
	d.tail = node
}

//	func (d *DoubleLinkedList[T]) InsertAt(item T, id int) {
//		if id > d.Length {
//			panic("error")
//		} else if id == d.Length {
//			d.Append(item)
//			return
//		} else if id == 0 {
//			d.Prepend(item)
//			return
//		}
//
//		d.Length++
//		curr := d.getAt(id)
//		node := newNode(item)
//		node.prev = curr
//		node.prev = curr.prev
//		curr.prev = node
//
//		if node.prev != nil {
//			node.prev.next = curr
//		}
//	}
func (d *DoubleLinkedList[T]) InsertAt(item T, id int) {
	if id > d.Length {
		panic("error")
	} else if id == d.Length {
		d.Append(item)
		return
	} else if id == 0 {
		d.Prepend(item)
		return
	}

	d.Length++
	curr := d.getAt(id)
	node := newNode(item)

	node.prev = curr.prev
	node.next = curr

	if curr.prev != nil {
		curr.prev.next = node
	}

	curr.prev = node
}

func (d *DoubleLinkedList[T]) Remove(item T) (T, error) {
	var zero T
	curr := d.head
	i := 0
	for curr != nil && i < d.Length {
		if isEqual(curr.value, item) {
			break
		}
		curr = curr.next
		i++ // Increment i to avoid infinite loop
	}

	if curr == nil {
		return zero, fmt.Errorf("invalid curr: %v", curr)
	}

	return d.removeNode(curr)
}

func (d *DoubleLinkedList[T]) Get(id int) (T, error) {
	var zero T
	node := d.getAt(id)
	if node == nil {
		return zero, fmt.Errorf("invalid node at position: %d", id)
	}
	return node.value, nil
}

func (d *DoubleLinkedList[T]) RemoveAt(id int) (T, error) {
	node := d.getAt(id)
	var zero T
	if node == nil {
		return zero, fmt.Errorf("invalid position: %d", id)
	}

	if id < 0 || id >= d.Length {
		return zero, fmt.Errorf("invalid position: %d", id)
	}
	return d.removeNode(node)
}

// private methods
func (d *DoubleLinkedList[T]) removeNode(node *LinkedListNode[T]) (T, error) {
	var zero T
	d.Length--
	if d.Length == 0 {
		if d.head != nil {
			out := d.head.value
			d.head, d.tail = nil, nil
			return out, fmt.Errorf("length: %d", d.Length)
		}
		return zero, fmt.Errorf("length: %d", d.Length)
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == d.head {
		d.head = node.next
	}

	if node == d.tail {
		d.tail = node.prev
	}
	node.prev, node.next = nil, nil
	return node.value, nil
}

func (d *DoubleLinkedList[T]) getAt(id int) *LinkedListNode[T] {
	curr := d.head
	i := 0
	for curr != nil && i < id {
		curr = curr.next
	}
	return curr
}
