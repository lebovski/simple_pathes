package graph

// Depth-first search for any values
func (edges *Edges) Dfs(from interface{}, to interface{}) [][]interface{} {
	uniqueVertexes := edges.GetUniqueVertexes()
	intToInterface, interfaceToInt := MakeInterfaceAndIntConverters(uniqueVertexes)
	intEdges := edges.ConvertEdgesToIntEdges(interfaceToInt)

	intFrom := (*interfaceToInt)[from]
	intTo := (*interfaceToInt)[to]

	intPaths := intEdges.IntDfs(intFrom, intTo)
	commonPaths := GetCommonPath(intPaths, intToInterface)
	return commonPaths
}
