package utils

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
	"github.com/lebovski/simple_pathes/graph"
)

type States map[string][]map[string][]map[string]string

type Action struct {
	Name string
	index int
}

func GetEdges(ymlPath string) *graph.Edges {
	var edges graph.Edges
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

	for stateIndex, v := range states["states"] {
		for state, actions := range v {
			for _, action := range actions {
				act := Action{Name:action["action"], index:stateIndex}
				edges = append(edges, graph.Edge{state, act})
				edges = append(edges, graph.Edge{act, action["result"]})
			}
		}
	}

	return &edges
}
