package main

import (
	"fmt"
	"strconv"
)

type TrNode struct {
	Val   string
	Left  *TrNode
	Right *TrNode
}

func NewTrNode(val string) *TrNode {
	return &TrNode{
		Val: val,
	}
}

func BuildTree(items [][3]string) *TrNode {
	nodes := make([]*TrNode, len(items))
	getNode := func(ind int) *TrNode {
		node := nodes[ind]
		if node == nil {
			node = NewTrNode("")
			nodes[ind] = node
		}
		return node
	}
	vt := make([]int8, len(items))
	for i, temp := range items {
		node := getNode(i)
		node.Val = temp[0]
		if temp[1] != "-" {
			ind, _ := strconv.Atoi(temp[1])
			node.Left = getNode(ind)
			vt[ind] = 1
		}
		if temp[2] != "-" {
			ind, _ := strconv.Atoi(temp[2])
			node.Right = getNode(ind)
			vt[ind] = 1
		}
	}
	root := 0
	for i, val := range vt {
		if val == 0 {
			root = i
			break
		}
	}
	return nodes[root]
}

func isOk(t1 *TrNode, t2 *TrNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	} else if t1.Val != t2.Val {
		return false
	}
	f1 := isOk(t1.Left, t2.Left) && isOk(t1.Right, t2.Right)
	f2 := isOk(t1.Right, t2.Left) && isOk(t1.Left, t2.Right)
	return f1 || f2
}

func main() {
	var n int
	nodes := [2]*TrNode{}
	for i := 0; i < 2; i++ {
		fmt.Scan(&n)
		var items [][3]string
		var a, b, c string
		for j := 0; j < n; j++ {
			fmt.Scan(&a, &b, &c)
			items = append(items, [3]string{a, b, c})
		}
		if n != 0 {
			nodes[i] = BuildTree(items)
		}
	}
	if isOk(nodes[0], nodes[1]) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
