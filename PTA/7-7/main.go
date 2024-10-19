package main

import (
	"bufio"
	"fmt"
	"os"
)

const N, M = 1001, 66002

var (
	H     [N]int
	E, NE [M]int
	ind   = 1
)

func add(a, b int) {
	E[ind], NE[ind], H[a] = b, H[a], ind
	ind++
}

func bfs(v int) int {
	var q [][2]int
	q = append(q, [2]int{v, 0})
	vt := make([]int, N)
	vt[v] = 1
	cnt := 0
	for len(q) > 0 {
		pos := q[0]
		cnt++
		if pos[1] < 6 {
			for i := H[pos[0]]; i != 0; i = NE[i] {
				w := E[i]
				if vt[w] == 0 {
					vt[w] = 1
					q = append(q, [2]int{w, pos[1] + 1})
				}
			}
		}
		q = q[1:]
	}
	return cnt
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m, a, b int
	fmt.Fscan(r, &n, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b)
		add(a, b)
		add(b, a)
	}
	for i := 0; i < n; i++ {
		cnt := bfs(i + 1)
		fmt.Printf("%d: %.2f%%\n", i+1, float64(cnt)/float64(n)*100)
	}
}
