package main

import (
	"bufio"
	"fmt"
	"os"
)

func merge_sort(arr []int, i, j int) {
	if i >= j {
		return
	}
	mid := (i + j) >> 1
	merge_sort(arr, i, mid)
	merge_sort(arr, mid+1, j)
	merge(arr, i, mid, j)
}

func merge(arr []int, l, mid, r int) {
	i, j, t := l, mid+1, 0
	temp := make([]int, r-l+1)
	for i <= mid && j <= r {
		if arr[i] <= arr[j] {
			temp[t] = arr[i]
			i++
		} else {
			temp[t] = arr[j]
			j++
		}
		t++
	}
	for i <= mid {
		temp[t] = arr[i]
		i++
		t++
	}
	for j <= mid {
		temp[t] = arr[j]
		j++
		t++
	}
	for i := 0; i < t; i++ {
		arr[i+l] = temp[i]
	}
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
		if i != 0 {
			fmt.Printf(" ")
		}
		fmt.Print(arr[i])
	}
}
