package main

import (
	"bufio"
	"fmt"
	"os"
)

var ant int = 1e9

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func dfs(arr [][]int, a, f, n int) int {
	val, sum := 0, 0
	for _, b := range arr[a] {
		if b != f {
			t := dfs(arr, b, a, n)
			val = max(val, t)
			sum += t
		}
	}
	ant = min(ant, max(n-sum-1, val))
	return sum + 1
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, a, b int
	fmt.Fscan(r, &n)
	arr := make([][]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(r, &a, &b)
		arr[a-1] = append(arr[a-1], b-1)
		arr[b-1] = append(arr[b-1], a-1)
	}
	dfs(arr, 0, -1, n)
	fmt.Println(ant)
}
