package day_graphs

type WeightedAdjacencyMatrix struct {
	Values [][]int
}

func BFSMatrixList(graph WeightedAdjacencyMatrix, source int, needle int) []int {
	seen := make([]bool, len(graph.Values))
	prev := make([]int, len(graph.Values))
	for i := range prev {
		prev[i] = -1
	}

	seen[source] = true
	q := []int{source}
	//check
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr == needle {
			break
		}

		aux := graph.Values[curr]
		for i := 0; i < len(aux); i++ {
			if aux[i] == 0 {
				continue
			}

			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr
			q = append(q, i)
		}
	}
	//build it backwards
	if prev[needle] == -1 {
		return []int{} // Return an empty integer slice
	}

	curr := needle
	var out []int
	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	result := []int{source}
	reversed := make([]int, 0)

	// Reverse the 'out' slice and store it in 'reversed'
	for i := len(out) - 1; i >= 0; i-- {
		reversed = append(reversed, out[i])
	}

	// Concatenate 'result' and 'reversed'
	result = append(result, reversed...)

	return result

}
