package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func maxSum2(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	cnt := len(arr) / 2
	lSum := maxSum2(arr[:cnt])
	rSum := maxSum2(arr[cnt:])
	a1, a2, t1, t2 := 0, 0, 0, 0
	for i := cnt - 1; i >= 0; i-- {
		t1 += arr[i]
		if t1 > a1 {
			a1 = t1
		}
	}
	for j := cnt; j < len(arr); j++ {
		t2 += arr[j]
		if t2 > a2 {
			a2 = t2
		}
	}
	return slices.Max([]int{lSum, rSum, a1 + a2})
}

func maxSum1(arr []int) int {
	ant, temp := 0, 0
	for i := range arr {
		temp += arr[i]
		if temp < 0 {
			temp = 0
		}
		if temp > ant {
			ant = temp
		}
	}
	return ant
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	fmt.Println(maxSum2(arr))
}
