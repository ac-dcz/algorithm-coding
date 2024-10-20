package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 0x3f3f3f3f

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func floyd(graph [][]int, n int) {
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				graph[i][j] = min(graph[i][j], graph[i][k]+graph[k][j])
			}
		}
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n, &m)
	arr := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			arr[i][j] = INF
		}
		arr[i][i] = 0
	}
	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b, &c)
		arr[a][b], arr[b][a] = c, c
	}
	floyd(arr, n)
	v, d := 0, INF
	for i := 1; i <= n; i++ {
		m := 0
		for j := 1; j <= n; j++ {
			m = max(arr[i][j], m)
		}
		if m < d {
			v, d = i, m
		}
	}
	fmt.Print(v)
	if v != 0 {
		fmt.Print(" ", d)
	}
}
