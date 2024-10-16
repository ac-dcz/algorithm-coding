package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
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
	p.data[e] += c
}

func (p *Ploy) Display() string {
	p.Zip()
	var items [][2]int
	for e, c := range p.data {
		items = append(items, [2]int{e, c})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] > items[j][0]
	})
	var words []string
	for _, item := range items {
		words = append(words, fmt.Sprintf("%d %d", item[1], item[0]))
	}
	if len(words) == 0 {
		return "0 0"
	}
	return strings.Join(words, " ")
}

func (p *Ploy) Zip() *Ploy {
	var items []int
	for e, c := range p.data {
		if c == 0 {
			items = append(items, e)
		}
	}
	for _, e := range items {
		delete(p.data, e)
	}
	return p
}

func (p *Ploy) Add(p1 *Ploy) *Ploy {
	p2 := NewPloy()
	for e, c := range p1.data {
		p2.data[e] += c
	}
	for e, c := range p.data {
		p2.data[e] += c
	}

	return p2.Zip()
}

func (p *Ploy) Mul(p1 *Ploy) *Ploy {
	p2 := NewPloy()
	for e1, c1 := range p.data {
		for e2, c2 := range p1.data {
			p2.Set(e1+e2, c1*c2)
		}
	}
	return p2.Zip()
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
	fmt.Println(p1.Mul(p2).Display())
	fmt.Println(p1.Add(p2).Display())
}
