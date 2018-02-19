package math

// Make adjacency directed matrix
func MakeAdjacencyDirectedMatrix(length int, edges IntEdges) [][]bool {
	squareMatrix := makeSquareMatrix(length)
	for _, edge := range edges {
		squareMatrix[edge[0]][edge[1]] = true
	}
	return squareMatrix
}

// Make adjacency undirected matrix
func MakeAdjacencyUndirectedMatrix(length int, edges IntEdges) [][]bool {
	squareMatrix := makeSquareMatrix(length)
	for _, edge := range edges {
		squareMatrix[edge[0]][edge[1]] = true
		squareMatrix[edge[1]][edge[0]] = true
	}
	return squareMatrix
}

// Make square matrix (length * length) and fill false values
func makeSquareMatrix(length int) [][]bool {
	var squareMatrix = make([][]bool, length)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			squareMatrix[i] = append(squareMatrix[i], false)
		}
	}
	return squareMatrix
}
