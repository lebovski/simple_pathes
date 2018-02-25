package utils

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"gopkg.in/yaml.v2"
	"github.com/lebovski/simple_pathes/graph"
)

type States map[string][]map[string][]map[string]string

func GetConfig() {
	var ces graph.Edges
	filename, _ := filepath.Abs("./utils/example.yml")

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	m := make(States)
	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		panic(err)
	}

	for u, v := range m["states"] {
		fff := strconv.Itoa(u)
		for i, j := range v {
			for _, c := range j {
				ces = append(ces, graph.Edge{i, [2]string{c["action"], fff}})
				ces = append(ces, graph.Edge{[2]string{c["action"], fff}, c["result"]})
			}
		}
	}
	res := ces.DFS("INITIAL", "BROKEN")
	fmt.Printf("\n%v ", len(res))
	for _, v := range res {
		fmt.Printf("\n%v ", v)
	}
	fmt.Printf("\n%v ", ces)
}
