package day_tree_search

import "github.com/FACorreiaa/prime-algo-go/day_tree"

func walkPost(curr *day_tree.BinaryNode[int], path []int) []int {
	if curr == nil {
		return path
	}

	//pre
	//recurse
	walkPost(curr.Left, path)
	walkPost(curr.Right, path)
	//post
	path = append(path, curr.Value)
	return path
}

func PostOrderSearch(head *day_tree.BinaryNode[int]) []int {
	var path []int
	walkPost(head, path)
	return path
}
