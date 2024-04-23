package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	vt, q := make([]int, n), make([][2]int, 0)
	q, vt[0] = append(q, [2]int{0, 0}), 1

	for len(q) > 0 {
		a, d := q[0][0], q[0][1]
		if a == n-1 {
			fmt.Println(d)
			break
		}
		for _, b := range arr[a] {
			if vt[b] == 0 {
				vt[b] = 1
				q = append(q, [2]int{b, d + 1})
			}
		}
		q = q[1:]
	}
	if vt[n-1] == 0 {
		fmt.Println(-1)
	}
}
