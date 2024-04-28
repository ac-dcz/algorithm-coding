package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF int = 0x7f7f7f7f
	N   int = 1e5
)

var (
	h, ne, w, nv     = make([]int, N), make([]int, N), make([]int, N), make([]int, N)
	ind          int = 0
)

func add(a, b, c int) {
	ne[ind], w[ind], nv[ind] = b, c, h[a]
	h[a] = ind
	ind++
}

func main() {
	for i := 0; i < N; i++ {
		h[i] = -1
	}
	r := bufio.NewReader(os.Stdin)
	var n, m, a, b, c int
	fmt.Fscan(r, &n, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b, &c)
		add(a-1, b-1, c)
	}
	dist, vt := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	dist[0] = 0
	q := []int{0}
	for len(q) > 0 {
		v := q[0]
		vt[v] = 0
		for i := h[v]; i != -1; i = nv[i] {
			if dist[v]+w[i] < dist[ne[i]] {
				dist[ne[i]] = dist[v] + w[i]
				if vt[ne[i]] == 0 {
					q = append(q, ne[i])
					vt[ne[i]] = 1
				}
			}
		}
		q = q[1:]
	}

	if dist[n-1] == INF {
		fmt.Println("impossible")
	} else {
		fmt.Println(dist[n-1])
	}
}
