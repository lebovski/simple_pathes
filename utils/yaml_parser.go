package utils

import (
	"io/ioutil"
	"path/filepath"

	"github.com/lebovski/simple_pathes/graph"
	"gopkg.in/yaml.v2"
)

// Get states from yaml file, for example see example.yml
func GetStatesFromYaml(ymlPath string) graph.States {
	filename, _ := filepath.Abs(ymlPath)

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	states := make(graph.States)
	err = yaml.Unmarshal(yamlFile, &states)
	if err != nil {
		panic(err)
	}

	return states
}
