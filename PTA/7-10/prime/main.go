package main

import (
	"bufio"
	"fmt"
	"os"
)

const N, M, INF = 1001, 3003, 0x3f3f3f3f

var (
	H        [N]int
	E, NE, W [M]int
	ind      = 1
)

func add(a, b, c int) {
	E[ind], NE[ind], W[ind], H[a] = b, H[a], c, ind
	ind++
}

func prime(n int) int {
	dist := make([]int, n+1)
	vt := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = INF
	}
	dist[1] = 0
	cost := 0
	for i := 0; i < n; i++ {
		t := -1
		for j := 1; j <= n; j++ {
			if vt[j] == 0 && (t == -1 || dist[j] < dist[t]) {
				t = j
			}
		}
		if dist[t] == INF {
			return -1
		}
		vt[t] = 1
		cost += dist[t]
		for j := H[t]; j != 0; j = NE[j] {
			b, c := E[j], W[j]
			if c < dist[b] {
				dist[b] = c
			}
		}
	}
	return cost
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n, &m)
	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b, &c)
		add(a, b, c)
		add(b, a, c)
	}
	fmt.Println(prime(n))
}
