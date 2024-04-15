package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	arr := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i][0], &arr[i][1])
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	cnt, right := 0, -1
	for i := 0; i < n; i++ {
		if arr[i][0] > right {
			cnt++
		}
		right = max(right, arr[i][1])
	}
	fmt.Println(cnt)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
