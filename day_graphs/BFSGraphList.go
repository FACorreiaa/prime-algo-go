package day_graphs

type GraphEdge struct {
	To     int
	Weight int
}

type WeightedAdjacencyList [][]GraphEdge

func walkBFSGraphList(graph WeightedAdjacencyList, curr int, needle int, seen []bool, path []int) bool {
	if seen[curr] {
		return false
	}
	seen[curr] = true

	//recurse pre
	if curr == needle {
		return true
	}

	//recurse
	list := graph[curr]
	for index, _ := range list {
		edge := list[index]
		if walkBFSGraphList(graph, edge.To, needle, seen, path) {
			path = append(path, edge.To)
			return true
		}
	}

	//post
	path = path[:len(path)-1]
	return false
}

func BFSGraphList(graph WeightedAdjacencyList, source int, needle int) ([]int, bool) {
	seen := make([]bool, len(graph))
	path := make([]int, 0)

	walkBFSGraphList(graph, source, needle, seen, path)
	if len(path) == 0 {
		return path, false
	}
	return path, true
}
