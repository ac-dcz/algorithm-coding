package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m, x, y, val int
	fmt.Fscan(r, &n, &m)
	nums, index := make([]int, 0), make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &x, &val)
		if _, ok := index[x]; ok {
			index[x] += val
		} else {
			index[x] = val
			nums = append(nums, x)
		}
	}
	index[int(-1e9-1)], index[int(1e9+1)] = 0, 0
	nums = append(nums, -1e9-1, 1e9+1)
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		index[nums[i]] += index[nums[i-1]]
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &x, &y)
		le := sort.SearchInts(nums, x)
		re := sort.SearchInts(nums, y)
		if nums[re] == y {
			re++
		}
		fmt.Println(index[nums[re-1]] - index[nums[le-1]])
	}

}
