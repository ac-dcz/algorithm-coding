package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)
	v, w := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &v[i], &w[i])
	}
	dp := make([]int, m+1)
	for i := 0; i < n; i++ {
		for j := m; j >= v[i]; j-- {
			if dp[j] < dp[j-v[i]]+w[i] {
				dp[j] = dp[j-v[i]] + w[i]
			}
		}
	}
	fmt.Println(dp[m])
}
