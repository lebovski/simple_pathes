package graph

// Маторица смежности
type graph [][]bool
//type pathes [][]int

// v - откуда начинаем
// x - где заканчиваем
func dfs(graph graph, visited []bool, n int, v int, x int, cnt int) int {
	if v == x {
		cnt++
		return cnt
	}
	visited[v] = true
	for i := 0; i <= n; i++ {
		if graph[v][i] && !visited[i] {
			cnt = dfs(graph, visited, n, i, x, cnt)
		}
	}
	visited[v] = false
	return cnt
}
