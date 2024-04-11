package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m, t int
	fmt.Fscan(r, &n, &m, &t)
	a, b := make([]int, n), make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}
	for j := 0; j < m; j++ {
		fmt.Fscan(r, &b[j])
	}
	i, j := 0, m-1
	for i < n && j >= 0 {
		val := a[i] + b[j]
		if val < t {
			i++
		} else if val > t {
			j--
		} else {
			break
		}
	}
	fmt.Println(i, j)
}
