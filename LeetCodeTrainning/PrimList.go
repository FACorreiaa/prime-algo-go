package leet_code_trainning

import (
	"github.com/FACorreiaa/prime-algo-go/day_graphs"
	"math"
)

// Prims algorithm :: minimum spanning tree
// What is a minimum spanning tree?
//  - Requires no cycles
//  - For it to technically be a minimum spanning tree, the graph requires
//  to be strongly connected
// 1. Select starting node,
// 2. put edges of current selected node into a list
// 3. select edge that is the lowest value and to a node we haven't seen yet
// 4. we need to insert the edge from current to new into our mst
// 5. the newly selected node becomes the current node,
// 6. repeat to step 2 until unvisited is empty or unreachable
// ...
// 8. $$

type EdgeTuple struct {
	First  int
	Second []day_graphs.GraphEdge
}

func PrimsList(list day_graphs.WeightedAdjacencyList) day_graphs.WeightedAdjacencyList {
	visited := make([]bool, len(list))
	mst := make(day_graphs.WeightedAdjacencyList, len(list))

	visited[0] = true
	current := 0

	edges := []EdgeTuple{
		{First: 0, Second: []day_graphs.GraphEdge{}},
	}
	for {
		// Put all edges from the current node into the list
		for _, edge := range list[current] {
			edges[0].Second = append(edges[0].Second, edge)
		}

		lowest := math.Inf(1) // or any initial value that represents positive infinity
		var lowestEdge EdgeTuple

		for _, edge := range edges {
			if len(edge.Second) == 0 {
				continue
			}

			target := edge.Second[0].To
			weight := float64(edge.Second[0].Weight)

			if !visited[target] && weight < lowest {
				lowest = weight
				lowestEdge = edge
			}
		}

		// Insert the edge into the MST and update visited
		if lowestEdge.Second != nil {
			from := lowestEdge.First
			to := lowestEdge.Second[0].To
			weight := lowestEdge.Second[0].Weight

			mst[from] = append(mst[from], day_graphs.GraphEdge{To: to, Weight: weight})
			mst[to] = append(mst[to], day_graphs.GraphEdge{To: from, Weight: weight})
			visited[to] = true

			// Remove the used edge from the list
			for i, edge := range edges {
				if edge.First == lowestEdge.First && len(edge.Second) == len(lowestEdge.Second) {
					equal := true
					for j := range edge.Second {
						if edge.Second[j] != lowestEdge.Second[j] {
							equal = false
							break
						}
					}
					if equal {
						edges = append(edges[:i], edges[i+1:]...)
						break
					}
				}
			}

		}

		// Find the newly selected node as the current node
		if lowestEdge.Second != nil {
			current = lowestEdge.Second[0].To
		} else {
			break
		}

		// Check if all nodes have been visited
		unvisited := false
		for _, v := range visited {
			if !v {
				unvisited = true
				break
			}
		}

		if !unvisited {
			break
		}
	}

	return mst
}
