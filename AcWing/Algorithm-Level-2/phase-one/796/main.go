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
			arr[i][j] += arr[i-1][j] + arr[i][j-1] - arr[i-1][j-1]
		}
	}
	var x1, y1, x2, y2 int
	for q > 0 {
		fmt.Fscan(r, &x1, &y1, &x2, &y2)
		fmt.Println(arr[x2][y2] - arr[x2][y1-1] - arr[x1-1][y2] + arr[x1-1][y1-1])
		q--
	}
}
