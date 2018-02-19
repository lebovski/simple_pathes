package graph

// Depth-first search for any values
func (edges *Edges) DFS(from interface{}, to interface{}) [][]interface{} {
	uniqueVertexes := edges.GetUniqueVertexes()
	intToInterface, interfaceToInt := MakeInterfaceAndIntConverters(uniqueVertexes)
	intEdges := edges.ConvertEdgesToIntEdges(interfaceToInt)

	intFrom := (*interfaceToInt)[from]
	intTo := (*interfaceToInt)[to]

	intPaths := intEdges.IntDFS(intFrom, intTo)
	commonPaths := GetCommonPath(intPaths, intToInterface)
	return commonPaths
}
