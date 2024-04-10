package main

import (
	"bytes"
	"fmt"
)

func Add(a, b string) string {
	c := bytes.NewBuffer(nil)
	i, j, r := len(a)-1, len(b)-1, 0
	for i >= 0 || j >= 0 || r != 0 {
		t := r
		if i >= 0 {
			t += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			t += int(b[j] - '0')
			j--
		}
		c.WriteByte(byte(t%10) + '0')
		r = t / 10
	}
	data := c.Bytes()
	// slices.Reverse[[]byte, byte](data)
	// slices.Reverse(data)
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return string(data)
}

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	c := Add(a, b)
	fmt.Println(c)
}
