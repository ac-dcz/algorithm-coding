package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type PolyMulity map[int]int

func NewPolyMulity(c, e []int) PolyMulity {
	p := make(PolyMulity)
	for i := range c {
		p[e[i]] = c[i]
	}
	return p
}

func (p PolyMulity) Add(q PolyMulity) PolyMulity {
	r := make(PolyMulity)
	for e, c := range p {
		r[e] = c
	}
	for e, c := range q {
		if _, ok := r[e]; ok {
			r[e] += c
		} else {
			r[e] = c
		}
		if r[e] == 0 {
			delete(r, e)
		}
	}
	return r
}

func (p PolyMulity) Mul(q PolyMulity) PolyMulity {
	r := make(PolyMulity)
	for e1, c1 := range q {
		for e2, c2 := range p {
			e, c := e1+e2, c1*c2
			if _, ok := r[e]; ok {
				r[e] += c
			} else {
				r[e] = c
			}
			if r[e] == 0 {
				delete(r, e)
			}
		}
	}
	return r
}

func (p PolyMulity) String() string {
	keys := make([]int, 0)
	for k := range p {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] >= keys[j]
	})
	items := make([]string, 0)
	for _, val := range keys {
		items = append(items, fmt.Sprintf("%d %d", p[val], val))
	}
	if len(items) == 0 {
		return "0 0"
	}
	return strings.Join(items, " ")
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	c1, e1 := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &c1[i], &e1[i])
	}
	fmt.Fscan(r, &n)
	c2, e2 := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &c2[i], &e2[i])
	}
	p1, p2 := NewPolyMulity(c1, e1), NewPolyMulity(c2, e2)
	fmt.Println(p1.Mul(p2))
	fmt.Println(p1.Add(p2))
}
