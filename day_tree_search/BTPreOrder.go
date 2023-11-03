package day_tree_search

import "github.com/FACorreiaa/prime-algo-go/day_tree"

func walkPre(curr *day_tree.BinaryNode[int], path []int) []int {
	if curr == nil {
		return path
	}

	//pre
	path = append(path, curr.Value)

	//recurse
	walkPre(curr.Left, path)
	walkPre(curr.Right, path)
	//post
	return path
}

func PreOrderSearch(head *day_tree.BinaryNode[int]) []int {
	var path []int
	walkPre(head, path)
	return path
}
