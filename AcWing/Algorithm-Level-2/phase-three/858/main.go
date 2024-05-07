package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   int = 500
	M   int = 2e5
	INF int = 0x7f7f7f7f
)

var (
	h   = make([]int, N)
	e   = make([]int, M)
	w   = make([]int, M)
	ne  = make([]int, M)
	ind = 0
)

func add(a, b, c int) {
	e[ind], w[ind], ne[ind] = b, c, h[a]
	h[a] = ind
	ind++
}

func prim(n int) int {
	ret, dist, vt := 0, make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	dist[0] = 0
	for i := 0; i < n; i++ {
		t := -1
		for j := 0; j < n; j++ {
			if vt[j] == 0 && (t == -1 || dist[j] < dist[t]) {
				t = j
			}
		}
		if dist[t] == INF {
			return INF
		}
		ret += dist[t]
		vt[t] = 1
		for j := h[t]; j != -1; j = ne[j] {
			if dist[e[j]] > w[j] {
				dist[e[j]] = w[j]
			}
		}
	}
	return ret
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
		add(b-1, a-1, c)
	}
	ret := prim(n)
	if ret == INF {
		fmt.Println("impossible")
	} else {
		fmt.Println(ret)
	}
}
