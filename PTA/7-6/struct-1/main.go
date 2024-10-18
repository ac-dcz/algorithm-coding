package main

import "fmt"

const N, M = 10, 100

var (
	H   [N]int
	E   [M]int
	NE  [M]int
	ind int = 0
)

func add(a, b int) {
	E[ind], NE[ind], H[a] = b, H[a], ind
	ind++
}

func dfs(v int, vt []int) []int {
	vt[v] = 1
	var path []int
	path = append(path, v)
	for i := H[v]; v != -1; i = NE[i] {
		if w := E[i]; vt[w] == 0 {
			path = append(path, dfs(w, vt)...)
		}
	}
	return path
}

func main() {
	for i := range H {
		H[i] = -1
	}
	var n, m int
	fmt.Scan(&n, &m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		add(a, b)
		add(b, a)
	}
}
