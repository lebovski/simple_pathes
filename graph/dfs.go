package graph

// v - откуда начинаем
// x - где заканчиваем
func dfs(graph [][]bool, visited []bool, n int, v int, x int, cnt int) int {
	visited[v] = true
	if v == x {
		cnt++
		return cnt
	}
	for i := 0; i <= n; i++ {
		if graph[v][i] && !visited[i] {
			cnt = dfs(graph, visited, n, i, x, cnt)
		}
	}
	visited[v] = false
	return cnt
}
