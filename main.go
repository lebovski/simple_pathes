package main

import (
	"fmt"

	"github.com/lebovski/simple_pathes/graph"
	"github.com/lebovski/simple_pathes/utils"
)

func main() {
	// CREATE SIMPLE GRAPH AND GET PATHS
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
	fmt.Println("\nGet vertexes from a to f")
	fmt.Printf("%v", "====== Simple paths")
	for _, v := range simpleEdgesRes {
		fmt.Printf("\n%v", v)
	}
	fmt.Printf("\n%v\n", "======")

	// CREATE GRAPH FROM YAML AND GET PATHS
	states := utils.GetStatesFromYaml("./example.yml")

	fmt.Println("\nGet actions from INITIAL to BROKEN")
	edges := graph.MakeEdges(states, 0)
	brokenRes := edges.DFS(
		graph.Vertex{Name: "INITIAL", Type: "state", Loop: false, LoopCount: 0, Index: 0},
		graph.Vertex{Name: "BROKEN", Type: "state", Loop: false, LoopCount: 0, Index: 0},
	)
	fmt.Printf("Broken count: %v\n%v", len(brokenRes), "====== Paths:")
	iniToBroPaths := graph.GetPathsOfNames(brokenRes, "Chunk")
	for _, v := range iniToBroPaths {
		fmt.Printf("\n%v", v)
	}
	fmt.Printf("\n%v\n", "======")

	fmt.Println("\nGet actions from INITIAL to FINISHED")
	finishedRes := edges.DFS(
		graph.Vertex{Name: "INITIAL", Type: "state", Loop: false, LoopCount: 0, Index: 0},
		graph.Vertex{Name: "FINISHED", Type: "state", Loop: false, LoopCount: 0, Index: 0},
	)

	fmt.Printf("Finished count: %v\n%v", len(finishedRes), "====== Paths:")
	iniToFinPaths := graph.GetPathsOfNames(finishedRes, "Chunk")
	for _, v := range iniToFinPaths {
		fmt.Printf("\n%v", v)
	}
	fmt.Printf("\n%v\n", "======")
}
