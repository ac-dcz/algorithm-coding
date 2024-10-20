package main

import (
	"bufio"
	"fmt"
	"os"
)

const N, M, INF = 101, 101 * 101, 0x3f3f3f3f

var (
	H        [N]int
	E, NE, W [M]int
	ind      int = 1
)

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

func add(a, b, c int) {
	E[ind], NE[ind], W[ind], H[a] = b, H[a], c, ind
	ind++
}

func dijkstra(v, n int) []int {
	vt := make([]int, n+1)
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = INF
	}
	dist[v] = 0
	for i := 0; i < n; i++ {
		t := -1
		for j := 1; j <= n; j++ {
			if dist[j] != INF && vt[j] == 0 && (t == -1 || dist[j] < dist[t]) {
				t = j
			}
		}
		if t == -1 {
			return nil
		}
		vt[t] = 1
		for j := H[t]; j != 0; j = NE[j] {
			e, w := E[j], W[j]
			dist[e] = min(dist[e], dist[t]+w)
		}
	}
	return dist
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m, a, b, c int
	fmt.Fscan(r, &n, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b, &c)
		add(a, b, c)
		add(b, a, c)
	}
	v, d := 0, INF
	for i := 1; i <= n; i++ {
		if dist := dijkstra(i, n); dist != nil {
			fmt.Println(dist)
			m := 0
			for _, val := range dist {
				m = max(m, val)
			}
			if m < d {
				v, d = i, m
			}
		}
	}
	if v == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(v, d)
	}
}
