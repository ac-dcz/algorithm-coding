package main

import (
	"bufio"
	"fmt"
	"os"
)

func quick_sort(arr []int, left, right int) {
	if right <= left {
		return
	}
	i, j, privot := left, right, arr[(left+right)>>1]
	for i < j {
		for arr[i] < privot {
			i++
		}
		for arr[j] > privot {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	quick_sort(arr, left, j)
	quick_sort(arr, i, right)
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
		fmt.Printf("%d ", arr[i])
	}
}
