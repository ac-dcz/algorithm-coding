package main

import (
	"bytes"
	"fmt"
)

func Mul(a string, b int) string {
	if b == 0 {
		return "0"
	}
	c := bytes.NewBuffer(nil)
	i, r := len(a)-1, 0
	for i >= 0 || r != 0 {
		if i >= 0 {
			r += int(a[i]-'0') * b
			i--
		}
		c.WriteByte(byte(r%10) + '0')
		r /= 10
	}
	data := c.Bytes()
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return c.String()
}

func main() {
	var (
		a string
		b int
	)
	fmt.Scan(&a, &b)
	c := Mul(a, b)
	fmt.Println(c)
}
