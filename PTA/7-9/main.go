package main

import (
	"bufio"
	"fmt"
	"os"
)

const N, M, INF = 500, 500000, 0x3f3f3f3f

var (
	H           [N]int
	E, NE, W, X [M]int
	ind         = 0
)

func add(a, b, c, d int) {
	E[ind], NE[ind], W[ind], X[ind], H[a] = b, H[a], c, d, ind
	ind++
}

func dijkstra(s, n int) ([]int, []int) {
	dist, pay := make([]int, n), make([]int, n)
	vt := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[s] = 0
	for i := 0; i < n; i++ {
		t := -1
		for j := 0; j < n; j++ {
			if vt[j] == 0 && (t == -1 || dist[j] < dist[t]) {
				t = j
			}
		}
		if dist[t] == INF {
			break
		}
		vt[t] = 1
		for j := H[t]; j != -1; j = NE[j] {
			b, c, d := E[j], W[j], X[j]
			if dist[b] > dist[t]+c {
				dist[b] = dist[t] + c
				pay[b] = pay[t] + d
			} else if dist[b] == dist[t]+c && pay[b] > pay[t]+d {
				pay[b] = pay[t] + d
			}
		}
	}
	return dist, pay
}

func main() {
	for i := range H {
		H[i] = -1
	}
	r := bufio.NewReader(os.Stdin)
	var n, m, s, d int
	fmt.Fscan(r, &n, &m, &s, &d)
	for i := 0; i < m; i++ {
		var a, b, c, d int
		fmt.Fscan(r, &a, &b, &c, &d)
		add(a, b, c, d)
		add(b, a, c, d)
	}
	dist, pay := dijkstra(s, n)
	fmt.Println(dist[d], pay[d])
}
