package day_tree

type BinaryNode[T any] struct {
	Value T
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
}

func Compare[T comparable](a *BinaryNode[T], b *BinaryNode[T]) bool {
	//structural check
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Value != b.Value {
		return false
	}

	return Compare(a.Left, b.Left) && Compare(a.Right, b.Right)
}
