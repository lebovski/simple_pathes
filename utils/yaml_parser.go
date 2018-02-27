package utils

import (
	"io/ioutil"
	"path/filepath"

	"github.com/lebovski/simple_pathes/graph"
	"gopkg.in/yaml.v2"
)

type States map[string][]map[string][]map[string]string

type Vertex struct {
	Name      string
	Type      string
	Loop      bool
	LoopCount int
	Index     int
}

func GetEdges(ymlPath string, loopCount int) graph.Edges {
	var edges graph.Edges

	states := unmarshalYaml(ymlPath)
	loopElements := getLoopElements(states)
	for stateIndex, v := range states["states"] {
		for state, actions := range v {
			for _, action := range actions {
				var st0, st1, act Vertex

				act = Vertex{Name: action["action"], Type: "action", Loop: false, LoopCount: 0, Index: stateIndex}
				st0 = Vertex{Name: state, Type: "state", Loop: loopElements[state], LoopCount: 0, Index: 0}
				st1 = Vertex{Name: action["result"], Type: "state", Loop: loopElements[action["result"]], LoopCount: 0, Index: 0}

				if state != action["result"] {
					edges = append(edges, graph.Edge{st0, act})
					edges = append(edges, graph.Edge{act, st1})
				}
			}
		}
	}

	appendCicledEdges(loopElements, &edges, loopCount)

	return edges
}

func getLoopElements(states States) map[string]bool {
	result := make(map[string]bool)
	for _, v := range states["states"] {
		for state, actions := range v {
			for _, action := range actions {
				if state == action["result"] {
					result[state] = true
				}
			}
		}
	}
	return result
}

func unmarshalYaml(ymlPath string) States {
	filename, _ := filepath.Abs(ymlPath)

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	states := make(States)
	err = yaml.Unmarshal(yamlFile, &states)
	if err != nil {
		panic(err)
	}

	return states
}

func appendCicledEdges(loopVertexes map[string]bool, edges *graph.Edges, loopCount int) {
	for vertexName := range loopVertexes {
		vertex := Vertex{Name: vertexName, Type: "state", Loop: true, LoopCount: 0, Index: 0}
		inEdges := getInEdges(*edges, vertex)
		outEdges := getOutEdges(*edges, vertex)

		for _, in := range inEdges {
			for _, out := range outEdges {
				for count := 0; count <= loopCount; count++ {
					if count == 0 {
						*edges = append(*edges, graph.Edge{in[0], out[1]})
					} else {
						newVertex := Vertex{Name: vertexName, Type: "state", Loop: true, LoopCount: count, Index: 0}
						*edges = append(*edges, graph.Edge{in[0], newVertex})
						*edges = append(*edges, graph.Edge{newVertex, out[1]})
					}
				}
			}
		}
	}
}

func getInEdges(edges graph.Edges, vertex Vertex) []graph.Edge {
	var res []graph.Edge
	for _, edge := range edges {
		if edge[1] == vertex {
			res = append(res, edge)
		}
	}
	return res
}

func getOutEdges(edges graph.Edges, vertex Vertex) []graph.Edge {
	var res []graph.Edge
	for _, edge := range edges {
		if edge[0] == vertex {
			res = append(res, edge)
		}
	}
	return res
}
