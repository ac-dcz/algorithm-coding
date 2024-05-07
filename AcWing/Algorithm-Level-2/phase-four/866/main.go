package main

import "fmt"

func Isprime(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n, num int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if Isprime(num) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
