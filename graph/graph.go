package graph

import (
	"fmt"
	"sort"
)

// Ребро
type Edge [2]int

// Список ребер
type Edges []Edge

// Получить все вершины
func (es *Edges) GetVertexes() []int {
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
	sort.Ints(vertexes)
	return vertexes
}

// Создать матрицу смежности
func (es *Edges) CreateAdjacencyMatrix(vertexes []int) [][]bool{
	var size = len(vertexes)
	var am = make([][]bool, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			am[i] = append(am[i], false)
		}
	}
	for _, v := range *es {
		am[v[0]][v[1]] = true
		am[v[1]][v[0]] = true
	}
	fmt.Printf("\n\n%v\n\n", am)
	return am
}

func Dfs(graph [][]bool, visited []bool, n int, v int, x int, cnt int) int {
	if v == x {
		cnt++
		fmt.Printf("\n%v\n", visited)
		return cnt
	}
	visited[v] = true
	for i := 0; i < n; i++ {
		if graph[v][i] && !visited[i] {
			cnt = Dfs(graph, visited, n, i, x, cnt)
		}
	}
	visited[v] = false
	return cnt
}