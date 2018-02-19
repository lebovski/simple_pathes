# simple_pathes

simple_pathes is library for get all pathes in graph using DFS.

Usage:

package main

```
import (
	"fmt"
	"github.com/lebovski/simple_pathes/graph"
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
}
```
