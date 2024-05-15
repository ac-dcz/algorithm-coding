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
	v, w := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &v[i], &w[i])
	}
	dp, ant := make([]int, m+1), 0
	for t := 1; t <= m; t++ {
		for i := 0; i < n; i++ {
			if v[i] <= t {
				if dp[t] < dp[t-v[i]]+w[i] {
					dp[t] = dp[t-v[i]] + w[i]
				}
			}
		}
		if dp[t] > ant {
			ant = dp[t]
		}
	}
	fmt.Println(ant)
}
