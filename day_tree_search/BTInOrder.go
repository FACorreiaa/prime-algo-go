package day_tree_search

import "github.com/FACorreiaa/prime-algo-go/day_tree"

func walkIn(curr *day_tree.BinaryNode[int], path []int) []int {
	if curr == nil {
		return path
	}

	//pre
	//recurse
	walkIn(curr.Left, path)
	path = append(path, curr.Value)
	walkIn(curr.Right, path)
	//post
	return path
}

func InOrderSearch(head *day_tree.BinaryNode[int]) []int {
	var path []int
	walkIn(head, path)
	return path
}
