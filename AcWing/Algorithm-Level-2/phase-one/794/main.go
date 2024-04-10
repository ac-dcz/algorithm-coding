package main

import (
	"bytes"
	"fmt"
)

func Div(a string, b int) (string, int) {
	c := bytes.NewBuffer(nil)
	r := 0
	for i := 0; i < len(a); i++ {
		r = r*10 + int(a[i]-'0')
		if r < b && c.Len() > 0 {
			c.WriteByte('0')
		} else if r >= b {
			c.WriteByte(byte(r/b) + '0')
		}
		r %= b
	}
	if c.Len() == 0 {
		c.WriteByte('0')
	}
	return c.String(), r
}

func main() {
	var (
		a string
		b int
	)
	fmt.Scan(&a, &b)
	c, r := Div(a, b)
	fmt.Println(c)
	fmt.Println(r)
}
