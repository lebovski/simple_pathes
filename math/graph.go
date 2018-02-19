package math

// Recursive depth-first search
func (edges *IntEdges) recursiveIntDFS(adjacencyMatrix *[][]bool, path []int, visited []bool, n int, v int, x int, results *[][]int) {
	path = append(path, v)
	if v == x {
		result := make([]int, len(path))
		copy(result, path[:])
		*results = append(*results, result)
		return
	}
	visited[v] = true
	for i := 0; i < n; i++ {
		if (*adjacencyMatrix)[v][i] && !visited[i] {
			edges.recursiveIntDFS(adjacencyMatrix, path, visited, n, i, x, results)
		}
	}
	visited[v] = false
	return
}

// Depth-first search for integer values
func (edges *IntEdges) IntDFS(from int, to int) *[][]int {
	vertexesCount := edges.getUniqueVertexesCount()
	adjacencyMatrix := MakeAdjacencyDirectedMatrix(vertexesCount, edges)

	visited := make([]bool, vertexesCount)
	results := make([][]int, 0)
	currentPath := make([]int, 0)

	edges.recursiveIntDFS(adjacencyMatrix, currentPath, visited, vertexesCount, from, to, &results)
	return &results
}
