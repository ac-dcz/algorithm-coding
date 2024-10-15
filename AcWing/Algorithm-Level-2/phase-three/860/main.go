package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1e5

var (
	h, e, ne, vt     = make([]int, N), make([]int, 2*N), make([]int, 2*N), make([]int, N)
	ind          int = 0
)

func add(a, b int) {
	e[ind], ne[ind] = b, h[a]
	h[a] = ind
	ind++
}

func dfs(v, c int, vt []int) bool {
	for i := h[v]; i != -1; i = ne[i] {
		b := e[i]
		if vt[b] == 1-c {
			return false
		} else if vt[b] == -1 {
			vt[b] = c
			if !dfs(b, 1-c, vt) {
				return false
			}
		}
	}
	return true
}

func main() {
	for i := range h {
		h[i] = -1
		vt[i] = -1
	}
	r := bufio.NewReader(os.Stdin)
	var n, m, a, b int
	fmt.Fscan(r, &n, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b)
		add(a-1, b-1)
		add(b-1, a-1)
	}

	flag := true
	for i := 0; i < n; i++ {
		if vt[i] == -1 {
			vt[i] = 0
			flag = dfs(i, 1, vt)
			if !flag {
				break
			}
		}
	}
	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
