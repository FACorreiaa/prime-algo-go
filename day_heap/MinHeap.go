package day_heap

import "math"

type MinHeap struct {
	Length int
	Data   []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		Length: 0,
		Data:   make([]int, 0),
	}
}

func (m *MinHeap) Insert(value int) {
	m.Data[m.Length] = value
	m.heapifyUp(m.Length)
	m.Length++
}

func (m *MinHeap) Delete() int {
	if m.Length == 0 {
		return -1
	}

	out := m.Data[0]
	m.Length--

	if m.Length == 0 {
		m.Data = make([]int, 0)
		return out
	}

	m.Data[0] = m.Data[m.Length]
	m.heapifyDown(0)
	return out
}

func (m *MinHeap) parent(id int) int {
	return int(math.Floor(float64(id-1) / 2))
}

func (m *MinHeap) leftChild(id int) int {
	return id*2 + 1
}

func (m *MinHeap) rightChild(id int) int {
	return id*2 + 2
}

func (m *MinHeap) heapifyDown(id int) {
	if id >= m.Length {
		return
	}

	leftId := m.leftChild(id)
	rightId := m.rightChild(id)

	if leftId >= m.Length || rightId >= m.Length {
		return
	}

	leftValue := m.Data[leftId]
	rightValue := m.Data[rightId]
	value := m.Data[id]

	if leftValue > rightValue && value > rightValue {
		m.Data[id] = rightValue
		m.Data[rightId] = value
		m.heapifyDown(rightId)
	} else if rightValue > leftValue && value > leftValue {
		m.Data[id] = leftValue
		m.Data[leftId] = value
		m.heapifyDown(leftId)
	}
}

func (m *MinHeap) heapifyUp(id int) {
	if id == 0 {
		return
	}

	p := m.parent(id)
	parentV := m.Data[p]
	v := m.Data[id]

	if parentV > v {
		m.Data[id] = parentV
		m.Data[p] = v
		m.heapifyUp(p)
	}
}
