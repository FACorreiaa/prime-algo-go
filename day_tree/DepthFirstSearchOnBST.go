package day_tree

func search[T comparable](curr *BinaryNode[T], needle T, less func(T, T) bool) bool {
	if curr == nil {
		return false
	}

	if curr.Value == needle {
		return true
	}

	if less(curr.Value, needle) {
		return search(curr.Right, needle, less)
	}

	return search(curr.Left, needle, less)
}

func DFS[T comparable](head *BinaryNode[T], needle T, less func(T, T) bool) bool {
	return search(head, needle, less)
}
