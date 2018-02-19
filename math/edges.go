package math

// List of graph edges
type IntEdges [][2]int

// Get count of unique vertexes
func (edges *IntEdges) getUniqueVertexesCount() int {
	var uniqueVertexes = make(map[int]bool, 0)
	for _, e := range *edges {
		for _, v := range e {
			if !uniqueVertexes[v] {
				uniqueVertexes[v] = true
			}
		}
	}
	return len(uniqueVertexes)
}
