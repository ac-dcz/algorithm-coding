package main

import (
	"bufio"
	"fmt"
	"os"
)

const N, M = 101, 10001

var (
	H, ES, LS [N]int
	E, NE, W  [M]int
	ind       = 1
)

func add(a, b, c int) {
	E[ind], NE[ind], W[ind], H[a] = b, H[a], c, ind
	ind++
}

func topological_sort(n int, in []int) (int, bool) {
	q, p := make([]int, 0), make([]int, 0)
	for i := 1; i <= n; i++ {
		if in[i] == 0 {
			q = append(q, i)
		}
	}
	t := 0
	for len(q) > 0 {
		a := q[0]
		p = append(p, a)
		for i := H[a]; i != 0; i = NE[i] {
			b, c := E[i], W[i]
			ES[b] = max(ES[b], ES[a]+c) //事件最早开始时间
			if in[b]--; in[b] == 0 {
				q = append(q, b)
			}
			t = max(t, ES[b])
		}
		q = q[1:]
	}
	if len(p) != n {
		return t, false
	}
	for len(p) > 0 {
		a := p[len(p)-1]
		LS[a] = t
		for i := H[a]; i != 0; i = NE[i] {
			b, c := E[i], W[i]
			LS[a] = min(LS[a], LS[b]-c) //事件的最晚开始时间
		}
		p = p[:len(p)-1]
	}

	return t, true
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n, &m)
	in := make([]int, n+1)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(r, &a, &b, &c)
		add(a, b, c)
		in[b]++
	}
	if t, ok := topological_sort(n, in); !ok {
		fmt.Println(0)
	} else {
		fmt.Println(t)
		for i := 1; i <= n; i++ {
			for j := H[i]; j != 0; j = NE[j] {
				b, c := E[j], W[j]
				if ES[i] == LS[b]-c {
					fmt.Printf("%d->%d\n", i, b)
				}
			}
		}
	}
}
