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
	pre  *Node
}

type List struct {
	head, tail *Node
	recored    []*Node
}

func NewList() *List {
	list := &List{
		head:    &Node{0, nil, nil},
		recored: make([]*Node, 1),
	}
	list.tail = list.head
	list.recored[0] = list.head
	return list
}

func (list *List) Insert(val, k, flag int) {
	node := list.tail
	if k != -1 {
		node = list.recored[k]
	}
	p := &Node{val, nil, nil}
	list.recored = append(list.recored, p)

	if flag == 0 {
		p.next = node
		p.pre = node.pre
		node.pre = p
		p.pre.next = p
	} else {
		p.next = node.next
		p.pre = node
		node.next = p
		p.next.pre = p
	}
	if list.tail.next != nil {
		list.tail = p
	}
}

func (list *List) Delete(k int) *Node {
	node := list.recored[k]
	pn := node.pre
	pn.next = node.next
	node.next.pre = pn
	return node
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
		case "L":
			fmt.Fscan(r, &val)
			list.Insert(val, 0, 1)
		case "R":
			fmt.Fscan(r, &val)
			list.Insert(val, -1, 1)
		case "D":
			fmt.Fscan(r, &k)
			list.Delete(k)
		case "IL":
			fmt.Fscan(r, &k, &val)
			list.Insert(val, k, 0)
		case "IR":
			fmt.Fscan(r, &k, &val)
			list.Insert(val, k, 1)
		}
		n--
	}
	fmt.Println(list)
}
