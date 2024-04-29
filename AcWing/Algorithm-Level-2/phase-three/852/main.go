package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   int = 2000
	M   int = 10000
	INF int = 0x7f7f7f7f
)

var (
	h, ne, w, nv     = make([]int, N), make([]int, M), make([]int, M), make([]int, M)
	ind          int = 0
)

func add(a, b, c int) {
	ne[ind], w[ind], nv[ind] = b, c, h[a]
	h[a] = ind
	ind++
}

func spfa(n int) bool {
	//不用特别处理dist
	dist, vt, pn, q := make([]int, n), make([]int, n), make([]int, n), make([]int, n) // pn[i]表示节点i最短路的长度
	for i := 0; i < n; i++ {
		q[i] = i //虚拟源节点
		vt[i] = 1
	}

	for len(q) > 0 {
		v := q[0]
		if pn[v] >= n {
			return true
		}
		vt[v] = 0
		for j := h[v]; j != -1; j = nv[j] {
			if dist[ne[j]] > dist[v]+w[j] {
				dist[ne[j]] = dist[v] + w[j]
				pn[ne[j]] = pn[v] + 1
				if vt[ne[j]] == 0 {
					q = append(q, ne[j])
					vt[ne[j]] = 1
				}
			}
		}
		q = q[1:]
	}
	return false
}

func main() {
	for i := range h {
		h[i] = -1
	}

	r := bufio.NewReader(os.Stdin)

	var n, m, a, b, c int

	fmt.Fscan(r, &n, &m)

	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b, &c)
		add(a-1, b-1, c)
	}

	flag := spfa(n)
	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
