package main

import (
	"bytes"
	"fmt"
)

func cmp(a, b string) int {
	if len(a) > len(b) || len(a) == len(b) && a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func Sub(a, b string) string {
	flag := cmp(a, b)
	if flag == 0 {
		return "0"
	} else if flag == -1 {
		a, b = b, a
	}
	c := bytes.NewBuffer(nil)
	i, j, r := len(a)-1, len(b)-1, 0
	for i >= 0 || j >= 0 || r != 0 {
		t := r
		if i >= 0 {
			t += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			t -= int(b[j] - '0')
			j--
		}
		c.WriteByte(byte((t+10)%10) + '0')
		r = 0
		if t < 0 {
			r = -1
		}
	}
	data := c.Bytes()
	// slices.Reverse[[]byte, byte](data)
	// slices.Reverse(data)
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	ind := 0
	for ind < len(data) && data[ind] == '0' {
		ind++
	}
	if flag == -1 {
		return "-" + string(data[ind:])
	}
	return string(data[ind:])
}

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	c := Sub(a, b)
	fmt.Println(c)
}
