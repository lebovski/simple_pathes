package main

import (
	"fmt"
	"github.com/lebovski/simple_pathes/graph"
	"github.com/lebovski/simple_pathes/utils"
)

func main() {
	var ces = graph.Edges{
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
	fmt.Printf("\n%v\n\n", ces.DFS("a", "f"))
	utils.GetConfig()
}
