package math

// Edge of graph
type IntEdge [2]int

// List of graph edges
type IntEdges []IntEdge

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
