package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m, q int
	fmt.Fscan(r, &n, &m, &q)
	arr := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(r, &arr[i][j])
		}
	}
	for i := n; i > 0; i-- {
		for j := m; j > 0; j-- {
			arr[i][j] -= arr[i][j-1] + arr[i-1][j] - arr[i-1][j-1]
		}
	}
	var x1, y1, x2, y2, val int
	for q > 0 {
		fmt.Fscan(r, &x1, &y1, &x2, &y2, &val)
		arr[x1][y1] += val
		if x2 < n {
			arr[x2+1][y1] -= val
		}
		if y2 < m {
			arr[x1][y2+1] -= val
		}
		if x2 < n && y2 < m {
			arr[x2+1][y2+1] += val
		}
		q--
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			arr[i][j] += arr[i][j-1] + arr[i-1][j] - arr[i-1][j-1]
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}
