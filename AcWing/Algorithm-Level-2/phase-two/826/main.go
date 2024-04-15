package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	val  int
	next *Node
}

type List struct {
	head    *Node
	recored []*Node
}

func NewList() *List {
	list := &List{
		head:    &Node{0, nil},
		recored: make([]*Node, 1),
	}
	list.recored[0] = list.head
	return list
}

func (list *List) Insert(val, k int) {
	node := list.recored[k]
	p := &Node{val, nil}
	list.recored = append(list.recored, p)
	p.next = node.next
	node.next = p
}

func (list *List) Delete(k int) *Node {
	node := list.recored[k]
	ret := node.next
	node.next = ret.next
	return ret
}

func (list *List) String() string {
	line, temp := strings.Builder{}, list.head
	for temp.next != nil {
		temp = temp.next
		line.WriteString(fmt.Sprintf("%d ", temp.val))
	}
	return line.String()
}

func main() {
	list := NewList()
	r := bufio.NewReader(os.Stdin)
	var (
		n, val, k int
		cmd       string
	)
	fmt.Fscan(r, &n)
	for n > 0 {
		fmt.Fscan(r, &cmd)
		switch cmd {
		case "H":
			fmt.Fscan(r, &val)
			list.Insert(val, 0)
		case "D":
			fmt.Fscan(r, &k)
			list.Delete(k)
		case "I":
			fmt.Fscan(r, &k, &val)
			list.Insert(val, k)
		}
		n--
	}
	fmt.Println(list)
}
