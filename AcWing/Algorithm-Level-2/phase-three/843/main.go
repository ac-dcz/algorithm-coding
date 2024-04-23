package main

import "fmt"

func dfs(arr, c, d, rd []int, ind, n int) {
	if ind == n {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if arr[i] == j {
					fmt.Printf("Q")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	} else {
		for i := 0; i < n; i++ {
			if c[i] == 1 || d[i+ind] == 1 || rd[n+ind-i-1] == 1 {
				continue
			}
			c[i], d[i+ind], rd[n+ind-i-1] = 1, 1, 1
			arr[ind] = i
			dfs(arr, c, d, rd, ind+1, n)
			c[i], d[i+ind], rd[n+ind-i-1] = 0, 0, 0
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	var (
		arr = make([]int, n)
		c   = make([]int, n)
		d   = make([]int, 2*n-1)
		rd  = make([]int, 2*n-1)
	)
	dfs(arr, c, d, rd, 0, n)
}
