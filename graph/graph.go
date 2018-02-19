package graph

// Depth-first search for any values
func (edges *Edges) DFS(from, to interface{}) [][]interface{} {
	uniqueVertexes := edges.GetUniqueVertexes()
	intToInterface, interfaceToInt := MakeInterfaceAndIntConverters(uniqueVertexes)
	intEdges := edges.ConvertEdgesToIntEdges(interfaceToInt)

	intFrom := interfaceToInt[from]
	intTo := interfaceToInt[to]

	intPaths := intEdges.IntDFS(intFrom, intTo)
	namedPath := GetNamedPath(intPaths, intToInterface)
	return namedPath
}
