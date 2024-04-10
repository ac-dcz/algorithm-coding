package main

import (
	"bufio"
	"fmt"
	"os"
)

// insert left
// func _LowerBound(arr []int, t int) int {
// 	l, r := 0, len(arr)-1
// 	for l <= r {
// 		mid := (l + r) >> 1
// 		if arr[mid] < t {
// 			l = mid + 1
// 		} else if arr[mid] >= t {
// 			r = mid - 1
// 		}
// 	}
// 	return l
// }
// insert right
// func _UpperBound(arr []int, t int) int {
// 	l, r := 0, len(arr)-1
// 	for l <= r {
// 		mid := (l + r) >> 1
// 		if arr[mid] <= t {
// 			l = mid + 1
// 		} else if arr[mid] > t {
// 			r = mid - 1
// 		}
// 	}
// 	return l
// }

func LowerBound(arr []int, t int) int {
	l, r, ret := 0, len(arr)-1, -1
	for l <= r {
		mid := (l + r) >> 1
		if arr[mid] < t {
			l = mid + 1
		} else if arr[mid] > t {
			r = mid - 1
		} else {
			ret = mid
			r = mid - 1
		}
	}
	return ret
}

func UpperBound(arr []int, t int) int {
	l, r, ret := 0, len(arr)-1, -1
	for l <= r {
		mid := (l + r) >> 1
		if arr[mid] <= t {
			l = mid + 1
		} else if arr[mid] > t {
			r = mid - 1
		} else {
			ret = mid
			l = mid + 1
		}
	}
	return ret
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, q int
	fmt.Fscan(r, &n, &q)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	var val int
	for q > 0 {
		fmt.Fscan(r, &val)
		l := LowerBound(arr, val)
		r := UpperBound(arr, val)
		fmt.Println(l, r)
		q--
	}
}
