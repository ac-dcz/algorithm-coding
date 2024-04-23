package main

import "fmt"

func dfs(arr []int, ind, state int) {
	if state == 0 {
		for _, val := range arr {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	} else {
		for i := 0; i < len(arr); i++ {
			if (state>>i)&1 == 1 {
				arr[ind] = i + 1
				dfs(arr, ind+1, state^(1<<i))
			}
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	dfs(arr, 0, (1<<n)-1)
}
