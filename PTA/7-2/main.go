package main

import (
	"bufio"
	"fmt"
	"os"
)

type Ploy struct {
	data map[int]int
}

func NewPloy() *Ploy {
	return &Ploy{
		data: make(map[int]int),
	}
}

func (p *Ploy) Set(e, c int) {
	p.data[e] = c
}

func (p *Ploy) Add(p1 *Ploy) *Ploy {
	p2 := NewPloy()
	for e, c := range p1.data {

	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, e, c int
	var (
		p1 = NewPloy()
		p2 = NewPloy()
	)
	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &c, &e)
		p1.Set(e, c)
	}
	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &c, &e)
		p2.Set(e, c)
	}

}
