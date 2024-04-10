package main

import (
	"fmt"
	"math"
)

func main() {
	var f float64
	fmt.Scan(&f)
	l, r := -100.0, 100.0
	for (r - l) >= (1e-8) {
		m := (l + r) / 2
		t := m * m * m
		if math.Abs((f - t)) < (1e-9) {
			break
		} else if f < t {
			r = m
		} else if f > t {
			l = m
		}
	}
	fmt.Printf("%.6f", (l+r)/2)
}
