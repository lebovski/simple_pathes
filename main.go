package main

import (
	"fmt"
	"github.com/lebovski/simple_pathes/graph"
	"github.com/lebovski/simple_pathes/utils"
	//"reflect"
)

func main() {
	var simpleEdges = graph.Edges{
		{"a", "b"},
		{"b", "f"},
		{"b", "d"},
		{"a", "c"},
		{"c", "d"},
		{"c", "j"},
		{"d", "e"},
		{"e", "f"},
		{"j", "f"},
		{"e", "c"},
	}
	simpleEdgesRes := simpleEdges.DFS("a", "f")
	fmt.Printf("\n%v", "====== Simple paths")
	for _, v := range simpleEdgesRes {
		fmt.Printf("\n%v", v)
	}

	edges := utils.GetEdges("./utils/example.yml")
	brokenRes := edges.DFS("INITIAL", "BROKEN")
	fmt.Printf("\n\nBroken count: %v", len(brokenRes))
	fmt.Printf("\n%v", "====== Paths:")
	for _, v := range brokenRes {
		fmt.Printf("\n%v", v)
	}

	finishedRes := edges.DFS("INITIAL", "FINISHED")
	fmt.Printf("\n\nFinished count: %v", len(finishedRes))
	fmt.Printf("\n%v", "====== Paths:")
	for _, v := range finishedRes {
		//for _, vv := range v {
		//	if reflect.TypeOf(vv) != reflect.TypeOf("") {
		//		fmt.Printf("\n%v", vv)
		//	}
		//}
		fmt.Printf("\n%v", v)
	}
}
