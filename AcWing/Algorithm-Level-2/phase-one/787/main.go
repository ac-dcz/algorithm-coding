package main

import (
	"bufio"
	"fmt"
	"os"
)

func merge_sort(arr []int, l, r int) {
	if l < r {
		mid := (l + r) >> 1
		merge_sort(arr, l, mid)
		merge_sort(arr, mid+1, r)
		merge(arr, l, mid, r)
	}
}

func merge(arr []int, l, le, re int) {
	i, j, L := l, le+1, 0
	temp := make([]int, re-l+1)
	for i <= le && j <= re {
		if arr[i] <= arr[j] {
			temp[L] = arr[i]
			i++
		} else {
			temp[L] = arr[j]
			j++
		}
		L++
	}
	for i <= le {
		temp[L] = arr[i]
		L++
		i++
	}
	for j <= re {
		temp[L] = arr[j]
		L++
		j++
	}
	copy(arr[l:re+1], temp)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	merge_sort(arr, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
