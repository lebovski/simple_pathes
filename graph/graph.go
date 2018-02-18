package graph

import (
	//"sort"
	//"fmt"
)

// Ребро
type edge [2]int

// Список ребер
type Edges []edge

// Получить все вершины
func (es *Edges) getVertexes() []int {
	var uniqueVertexes = make(map[int]bool)
	for _, e := range *es {
		for _, v := range e {
			if !uniqueVertexes[v] {
				uniqueVertexes[v] = true
			}
		}
	}
	var vertexes = make([]int, len(uniqueVertexes))
	for k := range uniqueVertexes {
		vertexes[k] = len(vertexes)
	}
	//sort.Ints(vertexes)
	return vertexes
}

// Создать матрицу смежности
func (es *Edges) createAdjacencyMatrix(vertexes []int) [][]bool{
	var size = len(vertexes)
	var am = make([][]bool, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			am[i] = append(am[i], false)
			//am[i] = append(am[i], false)
		}
	}
	//fmt.Printf("===\n%v\n===", am)
	for _, v := range *es {
		am[v[0]][v[1]] = true
		//am[v[1]][v[0]] = true
	}
	return am
}

// Рекурсивный поиск всех путей
func (es *Edges) recursiveDfs(graph [][]bool, path []int, visited []bool, n int, v int, x int, results *[][]int) {
	path = append(path, v)
	if v == x {
		result := make([]int, len(path))
		copy(result, path[:])
		*results = append(*results, result)
		return
	}
	visited[v] = true
	for i := 0; i < n; i++ {
		if graph[v][i] && !visited[i] {
			es.recursiveDfs(graph, path, visited, n, i, x, results)
		}
	}
	visited[v] = false
	return
}

// Поиск всех путей в графе вглубь
func (es *Edges) Dfs(from int, to int) [][]int {
	vertexes := es.getVertexes()
	adjacencyMatrix := es.createAdjacencyMatrix(vertexes)
	matrixLen := len(vertexes)

	visited := make([]bool, matrixLen)
	results := make([][]int, 0)
	currentPath := make([]int, 0)

	es.recursiveDfs(adjacencyMatrix, currentPath, visited, matrixLen, from, to, &results)
	return results
}