package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func find(parent []int, i int) int {
	if parent[i] != i {
		parent[i] = find(parent, parent[i])
		return parent[i]
	}
	return i
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n, &m)
	var edges [][3]int
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(r, &a, &b, &c)
		edges = append(edges, [3]int{a, b, c})
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	parent := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}
	cost, cnt := 0, 0
	for i := 0; i < m && cnt < n-1; i++ {
		edge := edges[i]
		a, b, c := edge[0], edge[1], edge[2]
		ra, rb := find(parent, a), find(parent, b)
		if ra != rb {
			parent[rb] = ra
			cost += c
			cnt++
		}
	}
	if cnt == n-1 {
		fmt.Println(cost)
	} else {
		fmt.Println(-1)
	}
}
