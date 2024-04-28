package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF int = 0x7f7f7f7f
	N   int = 500
	M   int = 1e4
)

var (
	h             = make([]int, N)
	ne, w, nv     = make([]int, M), make([]int, M), make([]int, M)
	ind       int = 0
)

func add(a, b, c int) {
	ne[ind], w[ind], nv[ind] = b, c, h[a]
	h[a] = ind
	ind++
}

func main() {
	for i := range h {
		h[i] = -1
	}
	r := bufio.NewReader(os.Stdin)
	var n, m, k, a, b, c int
	fmt.Fscan(r, &n, &m, &k)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b, &c)
		add(a-1, b-1, c)
	}
	dist, temp := make([]int, n), make([]int, n)
	for i := range dist {
		dist[i] = INF
	}

	dist[0] = 0
	for i := 0; i < k; i++ {
		copy(temp, dist)
		for v := 0; v < n; v++ {
			for j := h[v]; j != -1; j = nv[j] {
				if dist[ne[j]] > temp[v]+w[j] {
					dist[ne[j]] = temp[v] + w[j]
				}
			}
		}
	}

	if dist[n-1] >= INF/2 {
		fmt.Println("impossible")
	} else {
		fmt.Println(dist[n-1])
	}

	// dist[0] = 0
	// q := []int{0}
	// for k > 0 {
	// 	t := len(q)
	// 	vt := make([]int, n)
	// 	copy(temp, dist)
	// 	for i := 0; i < t; i++ {
	// 		v := q[0]
	// 		for j := h[v]; j != -1; j = nv[j] {
	// 			if dist[ne[j]] > temp[v]+w[j] {
	// 				dist[ne[j]] = temp[v] + w[j]
	// 				if vt[ne[j]] == 0 {
	// 					q = append(q, ne[j])
	// 					vt[ne[j]] = 1
	// 				}
	// 			}
	// 		}
	// 		q = q[1:]
	// 	}
	// 	k--
	// }
	// if dist[n-1] == INF {
	// 	fmt.Println("impossible")
	// } else {
	// 	fmt.Println(dist[n-1])
	// }
}
