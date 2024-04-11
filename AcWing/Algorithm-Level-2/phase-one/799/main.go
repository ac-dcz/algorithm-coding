package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	index := make(map[int]int)
	i, j, ant := 0, 0, 0
	for j < n {
		if val, ok := index[arr[j]]; ok && val != -1 {
			for ; i <= val; i++ {
				index[arr[i]] = -1
			}
		}
		index[arr[j]] = j
		if j-i+1 > ant {
			ant = j - i + 1
		}
		j++
	}
	fmt.Println(ant)
}
