package day_tree

//Breadth First Search

func newNode[T any](item T) *BinaryNode[T] {
	return &BinaryNode[T]{Value: item}
}

func BreadthFirstSearch[T comparable](head *BinaryNode[T], needle T) bool {
	if head == nil {
		return false
	}

	queue := []*BinaryNode[T]{head}
	//queue := newNode(root)
	for len(queue) > 0 {
		curr := queue[0]
		queue := queue[1:]

		// Search
		if curr != nil && curr.Value == needle {
			return true
		}

		// Enqueue left and right children
		if curr != nil {
			queue = append(queue, curr.Left, curr.Right)
		}
	}

	return false
}
