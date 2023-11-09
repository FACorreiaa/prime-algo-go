package day_graphs

import "math"

func hasUnvisited(seen []bool, dists []float64) bool {
	for i, s := range seen {
		if !s && dists[i] < math.Inf(1) {
			return true
		}
	}
	return false
}

func getLowestUnvisited(seen []bool, dists []float64) int {
	idx := -1
	lowestDistance := math.Inf(1)

	for i := 0; i < len(seen); i++ {
		if seen[i] {
			continue
		}

		if lowestDistance > dists[i] {
			lowestDistance = dists[i]
			idx = i
		}
	}
	return idx
}

func DijkstraList(source, sink int, arr WeightedAdjacencyList) []int {
	seen := make([]bool, len(arr))
	prev := make([]int, len(arr))
	dists := make([]float64, len(arr))

	for i := range dists {
		dists[i] = math.Inf(1)
	}
	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true

		aux := arr[curr]
		for i := 0; i < len(aux); i++ {
			edge := aux[i]
			if seen[edge.To] {
				continue
			}

			dist := dists[curr] + float64(edge.Weight)
			if dist < dists[edge.To] {
				dists[edge.To] = dist
				prev[edge.To] = curr
			}
		}
	}

	out := []int{sink} // Initialize the output slice with the sink node
	curr := sink
	for prev[curr] != source {
		curr = prev[curr]
		out = append([]int{curr}, out...) // Insert elements at the beginning
	}
	out = append([]int{source}, out...)

	return out
}
