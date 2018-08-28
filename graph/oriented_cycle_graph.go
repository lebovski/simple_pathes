package graph

// TODO: use structure
// `States` is map of vertexes which contains actions and result.
// Expected that action is edge which connect vertex (some state) and result (some state).
type States map[string][]map[string][]map[string]string

// Structure `Vertex` store info about vertex for work with loops
type Vertex struct {
	Name      string
	Type      string
	Loop      bool
	LoopCount int
	Index     int
}

// Get edges from states and use loop on elements with loops
func MakeEdges(states States, loopCount int) Edges {
	var edges Edges
	// Map of loop vertexes
	loopVertexes := getLoopVertexes(states)
	for stateIndex, v := range states[mainSection] {
		for state, actions := range v {
			for _, action := range actions {
				var beforeState, resultState, currentAction Vertex

				currentAction = Vertex{Name: action[actionVertexType], Type: actionVertexType, Loop: false, LoopCount: 0, Index: stateIndex}
				beforeState = Vertex{Name: state, Type: stateVertexType, Loop: loopVertexes[state], LoopCount: 0, Index: 0}
				resultState = Vertex{Name: action[resultOfAction], Type: stateVertexType, Loop: loopVertexes[action[resultOfAction]], LoopCount: 0, Index: 0}

				if state != action[resultOfAction] {
					edges = append(edges, Edge{beforeState, currentAction})
					edges = append(edges, Edge{currentAction, resultState})
				}
			}
		}
	}
	// Append loops
	appendCycledEdges(loopVertexes, &edges, loopCount)
	return edges
}

// Get list of actions lists something like: [[do_something_1, do_something_2, ...], [...] ...]
func GetPathsOfNames(paths [][]interface{}, actionOnLoop []string) [][]string {
	results := make([][]string, 0, len(paths))
	for _, path := range paths {
		actionsList := []string{"", }
		if hasPathLoops(path) {
			actionsList = actionOnLoop
		}
		for _, a := range actionsList {
			res := make([]string, 0, len(path))
			for _, ver := range path {
				vv := ver.(Vertex)
				if vv.Loop && vv.Type == stateVertexType {
					for ii := 0; ii <= vv.LoopCount; ii++ {
						res = append(res, a)
						continue
					}
				}
				if vv.Type == actionVertexType {
					res = append(res, vv.Name)
					continue
				}
			}
			results = append(results, res)
		}
	}
	return results
}

// Get vertexes with loops from states
func getLoopVertexes(states States) map[string]bool {
	result := make(map[string]bool)
	for _, v := range states[mainSection] {
		for state, actions := range v {
			for _, action := range actions {
				if state == action[resultOfAction] {
					result[state] = true
				}
			}
		}
	}
	return result
}

// Append cycles to list of edges
func appendCycledEdges(loopVertexes map[string]bool, edges *Edges, loopCount int) {
	for vertexName := range loopVertexes {
		vertex := Vertex{Name: vertexName, Type: stateVertexType, Loop: true, LoopCount: 0, Index: 0}
		inEdges := getInEdges(*edges, vertex)
		outEdges := getOutEdges(*edges, vertex)

		for _, in := range inEdges {
			for _, out := range outEdges {
				for count := 0; count <= loopCount; count++ {
					if count == 0 {
						// If we have Edges [a-b, b-c] then we create edge [a-c] i.e. without cycle
						*edges = append(*edges, Edge{in[0], out[1]})
					} else {
						// Create new virtual vertex for extend cycles
						newVertex := Vertex{Name: vertexName, Type: stateVertexType, Loop: true, LoopCount: count, Index: 0}
						*edges = append(*edges, Edge{in[0], newVertex})
						*edges = append(*edges, Edge{newVertex, out[1]})
					}
				}
			}
		}
	}
}

// Get edges which connect to vertex
func getInEdges(edges Edges, vertex Vertex) []Edge {
	var res []Edge
	for _, edge := range edges {
		if edge[1] == vertex {
			res = append(res, edge)
		}
	}
	return res
}

// Get edges which connect from vertex
func getOutEdges(edges Edges, vertex Vertex) []Edge {
	var res []Edge
	for _, edge := range edges {
		if edge[0] == vertex {
			res = append(res, edge)
		}
	}
	return res
}

func hasPathLoops(path []interface{}) bool {
	for _, ver := range path {
		vv := ver.(Vertex)
		if vv.Loop && vv.Type == stateVertexType {
			return true
		}
	}
	return false
}
