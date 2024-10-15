package main

import (
	"bufio"
	"fmt"
	"os"
)

// 递归方法
func recursionMethod(arr []int) int {
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	}
	mid := len(arr) / 2
	m1, m2 := recursionMethod(arr[:mid]), recursionMethod(arr[mid:])
	s1, s2 := 0, 0
	for i, s := mid-1, 0; i >= 0; i-- {
		s += arr[i]
		s1 = max(s1, s)
	}
	for i, s := mid, 0; i < len(arr); i++ {
		s += arr[i]
		s2 = max(s2, s)
	}
	return max(s1+s2, max(m1, m2))
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	// fmt.Println(max(0, recursionMethod(arr)))

	maxSum, curSum := 0, 0
	for i := 0; i < n; i++ {
		curSum += arr[i]
		if curSum < 0 {
			curSum = 0
		}
		maxSum = max(maxSum, curSum)
	}
	fmt.Println(maxSum)
}
