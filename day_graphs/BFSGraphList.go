package day_graphs

type GraphEdge struct {
	To     int
	Weight int
}

type WeightedAdjacencyList [][]GraphEdge

func walkBFSGraphList(graph WeightedAdjacencyList, curr int, needle int, seen []bool, path *[]int) bool {
	if seen[curr] {
		return false
	}
	seen[curr] = true

	// Append the current node to the path
	*path = append(*path, curr)

	//recurse pre
	if curr == needle {
		return true
	}

	//recurse
	//list := graph[curr]
	//for index, _ := range list {
	//	edge := list[index]
	//	if walkBFSGraphList(graph, edge.To, needle, seen, path) {
	//		*path = append(*path, edge.To)
	//		return true
	//	}
	//}
	list := graph[curr]
	for _, edge := range list {
		if walkBFSGraphList(graph, edge.To, needle, seen, path) {
			// No need to append edge.To here
			return true
		}
	}

	//post
	// If the current node is not part of the path, remove it
	if (*path)[len(*path)-1] == curr {
		*path = (*path)[:len(*path)-1]
	}

	return false
}

func BFSGraphList(graph WeightedAdjacencyList, source int, needle int) ([]int, bool) {
	seen := make([]bool, len(graph))
	path := make([]int, 0)

	walkBFSGraphList(graph, source, needle, seen, &path)
	if len(path) == 0 {
		return []int{}, false
	}
	return path, true
}
