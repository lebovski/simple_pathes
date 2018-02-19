package graph

import (
	"github.com/lebovski/simple_pathes/math"
)

// Named edge
type edge [2]interface{}

// List of named edges
type Edges []edge

// Get all unique vertexes from edges list
func (edges *Edges) GetUniqueVertexes() []interface{} {
	var uniqueVertexes = make(map[interface{}]bool)
	for _, e := range *edges {
		for _, v := range e {
			if !uniqueVertexes[v] {
				uniqueVertexes[v] = true
			}
		}
	}

	var vertexes = make([]interface{}, 0, len(uniqueVertexes))
	for uv := range uniqueVertexes {
		vertexes = append(vertexes, uv)
	}
	return vertexes
}

// Convert integer named edges to int edges
func (edges *Edges) ConvertEdgesToIntEdges(interfaceToInt map[interface{}]int) math.IntEdges {
	intEdges := make(math.IntEdges, 0, len(*edges))
	for _, edge := range *edges {
		intEdge := edge.ConvertEdgeToIntEdge(interfaceToInt)
		intEdges = append(intEdges, intEdge)
	}
	return intEdges
}

// Convert integer named edge to int edge
func (edge *edge) ConvertEdgeToIntEdge(interfaceToInt map[interface{}]int) [2]int {
	return [2]int{
		interfaceToInt[edge[0]], interfaceToInt[edge[1]],
	}
}
