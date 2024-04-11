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
	a, b := make([]int, n), make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &b[i])
	}
	i, j := 0, 0
	for i < n && j < m {
		if a[i] == b[j] {
			i++
		}
		j++
	}
	if i == n {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
