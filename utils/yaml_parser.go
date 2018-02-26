package utils

import (
	"io/ioutil"
	"path/filepath"

	"github.com/lebovski/simple_pathes/graph"
	"gopkg.in/yaml.v2"
)

type States map[string][]map[string][]map[string]string

type Vertex struct {
	Name  string
	Type  string
	Loop  bool
	Index int
}

func GetEdges(ymlPath string) graph.Edges {
	var edges graph.Edges

	states := unmarshalYaml(ymlPath)
	loopElements := getLoopElements(states)
	for stateIndex, v := range states["states"] {
		for state, actions := range v {
			for _, action := range actions {
				var st0, st1, act Vertex

				act = Vertex{Name: action["action"], Type: "action", Loop: false, Index: stateIndex}
				st0 = Vertex{Name: state, Type: "state", Loop: loopElements[state], Index: 0}
				st1 = Vertex{Name: action["result"], Type: "state", Loop: loopElements[action["result"]], Index: 0}

				edges = append(edges, graph.Edge{st0, act})
				edges = append(edges, graph.Edge{act, st1})
			}
		}
	}

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
