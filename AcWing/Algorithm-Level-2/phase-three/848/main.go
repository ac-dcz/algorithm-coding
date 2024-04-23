package main

import (
	"bufio"
	"fmt"
	"os"
)

func topSort(arr [][]int) []int {
	var q, t []int
	n := len(arr)
	indgree := make([]int, n)
	for i := 0; i < n; i++ {
		for _, v := range arr[i] {
			indgree[v]++
		}
	}
	for i, d := range indgree {
		if d == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		a := q[0]
		for _, b := range arr[a] {
			indgree[b]--
			if indgree[b] == 0 {
				q = append(q, b)
			}
		}
		t = append(t, a)
		q = q[1:]
	}
	return t
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n, &m)
	arr := make([][]int, n)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b)
		arr[a-1] = append(arr[a-1], b-1)
	}
	t := topSort(arr)
	if len(t) != n {
		fmt.Println(-1)
	} else {
		for _, v := range t {
			fmt.Printf("%d ", v+1)
		}
	}
}
