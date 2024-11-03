package main

import (
	"bufio"
	"fmt"
	"os"
)

func quick_sort(arr []int, i, j int) {
	if i >= j {
		return
	}
	l, r, v := i, j, arr[(i+j)>>1]
	for l < r {
		for arr[l] < v {
			l++
		}
		for arr[r] > v {
			r--
		}
		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	quick_sort(arr, i, r)
	quick_sort(arr, l, j)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	quick_sort(arr, 0, n-1)
	for i := 0; i < n; i++ {
		if i != 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", arr[i])
	}
}
