package main

import (
	"bufio"
	"fmt"
	"os"
)

func quick_sort(arr []int, left, right, k int) int {
	if right <= left {
		return arr[k]
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
	if k >= i {
		return quick_sort(arr, i, right, k)
	} else if k <= j {
		return quick_sort(arr, left, j, k)
	}
	return privot
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(r, &n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	val := quick_sort(arr, 0, n-1, k-1)
	fmt.Println(val)
}
