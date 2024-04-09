package main

import (
	"bufio"
	"fmt"
	"os"
)

func merge_sort(arr []int, l, r int) int {
	if l < r {
		mid := (l + r) >> 1
		ln := merge_sort(arr, l, mid)
		rn := merge_sort(arr, mid+1, r)
		mn := merge(arr, l, mid, r)
		return ln + rn + mn
	}
	return 0
}

func merge(arr []int, l, le, re int) int {
	i, j, t, ret := l, le+1, 0, 0
	temp := make([]int, re-l+1)
	for i <= le && j <= re {
		if arr[i] > arr[j] {
			ret += re - j + 1
			temp[t] = arr[i]
			i++
		} else {
			temp[t] = arr[j]
			j++
		}
		t++
	}
	for i <= le {
		temp[t] = arr[i]
		t++
		i++
	}
	for j <= re {
		temp[t] = arr[j]
		t++
		j++
	}
	copy(arr[l:re+1], temp)
	return ret
}

type MyErr string

func (m MyErr) String() string {
	return string(m)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	fmt.Println(merge_sort(arr, 0, n-1))
}
