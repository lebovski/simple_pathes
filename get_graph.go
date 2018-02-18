package main

import (
	"fmt"
	"simple_pathes/graph"
)

func main() {
	//var es = graph.Edges{
	//	{0, 1},
	//	{1, 2},
	//	{2, 3},
	//	{4, 5},
	//	{1, 7},
	//	{7, 4},
	//	{4, 2},
	//	{6, 7},
	//	{0, 7},
	//	{4, 8},
	//	{8, 3},
	//	{0, 9},
	//	{9, 3},
	//}
	//var es = graph.Edges{
	//	{0, 1},
	//	{1, 3},
	//	{0, 2},
	//	{2, 3},
	//	{1, 2},
	//}
	var ces = graph.CEdges{
		{"a", "b"},
		{"b", "d"},
		{"a", "c"},
		{"c", "d"},
		{"b", "c"},
	}
	aa := ces.GetUniqueVertexes()
	fmt.Printf("%v\n", ces.GetUniqueVertexes())
	intToInter, InterToInt := graph.MakeIntM(&aa)
	fmt.Printf("%v\n%v\n\n", intToInter, InterToInt)
	fmt.Printf("\n%v\n\n", ces.FromCEtoE(intToInter, InterToInt))
	gg := ces.FromCEtoE(intToInter, InterToInt)
	//paths := gg.Dfs(0, 3)
	//fmt.Printf("\n%v\n\n", paths)
	fmt.Printf("\n%v\n\n", graph.Dfs(&gg, "a", "d", InterToInt, intToInter))
	//res := es.Dfs(0, 3)
	//fmt.Printf("%v", res)
}