package graph

import (
	"github.com/lebovski/simple_pathes/math"
)

type edge [2]interface{}

type Edges []edge

// Get all unique vertexes from edges list
func (edges *Edges) GetUniqueVertexes() *[]interface{} {
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
	return &vertexes
}

// Convert integer path to common (interface) path
func (edges *Edges) ConvertEdgesToIntEdges(interfaceToInt *map[interface{}]int) *math.IntEdges {
	intEdges := make(math.IntEdges, 0, len(*edges))
	for _, v := range *edges {
		intEdges = append(intEdges, [2]int{(*interfaceToInt)[v[0]], (*interfaceToInt)[v[1]]})
	}
	return &intEdges
}
