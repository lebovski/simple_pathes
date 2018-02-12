package main

import (
	"fmt"
	"simple_pathes/graph"
)

func main() {
	var es = graph.Edges{
		{0, 1},
		{1, 2},
		{2, 3},
		{2, 4},
		{4, 5},
		{1, 7},
		{4, 7},
		{6, 7},
	}
	vv := es.GetVertexes()
	m := es.CreateAdjacencyMatrix(vv)
	vis := make([]bool, len(vv))
	res := graph.Dfs(m, vis, len(vv), 0, 3, 0)
	fmt.Printf("%v", res)
}