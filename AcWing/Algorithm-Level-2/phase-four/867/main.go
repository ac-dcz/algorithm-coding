package main

import "fmt"

func main() {
	var n, num int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		j := 2
		for num > 1 && j*j <= num {
			ret := 0
			for num%j == 0 {
				num /= j
				ret++
			}
			if ret > 0 {
				fmt.Println(j, ret)
			}
			j++
		}
		if num > 1 {
			fmt.Println(num, 1)
		}
		fmt.Println()
	}
}
