package main

import "fmt"

type Node struct {
	conf, expn int
	next       *Node
}

type Poly struct {
	nodes *Node
}

func NewNode(c, e int) *Node {
	node := new(Node)
	node.conf, node.expn = c, e
	node.next = nil
	return node
}

func NewPoly() *Poly {
	poly := new(Poly)
	poly.nodes = NewNode(10000, 10000)
	return poly
}

func (this *Node) Insert(node *Node) {
	node.next = this.next
	this.next = node
}

func (this *Poly) OrderInsert(node *Node) {
	temp := this.nodes
	for temp.next != nil {
		pn := temp.next
		if pn.expn < node.expn {
			temp.Insert(node)
			break
		} else if pn.expn == node.expn {
			pn.conf += node.conf
			break
		}
		temp = temp.next
	}
	if temp.next == nil {
		temp.Insert(node)
	}
}

func PolyAdd(p1, p2 *Poly) *Poly {
	p3 := NewPoly()
	thead := p3.nodes
	t1, t2 := p1.nodes.next, p2.nodes.next
	for t1 != nil || t2 != nil {
		if t1 != nil && t2 != nil {
			if t1.expn > t2.expn {
				thead.Insert(NewNode(t1.conf, t1.expn))
				t1 = t1.next
			} else if t1.expn < t2.expn {
				thead.Insert(NewNode(t2.conf, t2.expn))
				t2 = t2.next
			} else {
				thead.Insert(NewNode(t1.conf+t2.conf, t1.expn))
				t1, t2 = t1.next, t2.next
			}
		} else if t1 != nil {
			thead.Insert(NewNode(t1.conf, t1.expn))
			t1 = t1.next
		} else {
			thead.Insert(NewNode(t2.conf, t2.expn))
			t2 = t2.next
		}
		thead = thead.next
	}
	return p3
}

func PolyMutiply(p1, p2 *Poly) *Poly {
	p3 := NewPoly()
	t1 := p1.nodes.next

	for t1 != nil {
		t2 := p2.nodes.next
		for t2 != nil {
			pn := NewNode(t1.conf*t2.conf, t1.expn+t2.expn)
			p3.OrderInsert(pn)
			t2 = t2.next
		}
		t1 = t1.next
	}
	return p3
}

func (this *Poly) FormatOutPutPoly() {
	node := this.nodes.next
	cnt := 0
	for node != nil {
		if node.conf != 0 {
			if cnt != 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%d %d", node.conf, node.expn)
			cnt++
		}
		node = node.next
	}
	if cnt == 0 {
		fmt.Printf("0 0")
	}
	fmt.Printf("\n")
}

func main() {
	var n int
	p1, p2 := NewPoly(), NewPoly()
	fmt.Scan(&n)
	var c, e int
	t1 := p1.nodes
	for i := 0; i < n; i++ {
		fmt.Scan(&c, &e)
		pn := NewNode(c, e)
		t1.Insert(pn)
		t1 = t1.next
	}
	fmt.Scan(&n)
	t2 := p2.nodes
	for i := 0; i < n; i++ {
		fmt.Scan(&c, &e)
		pn := NewNode(c, e)
		t2.Insert(pn)
		t2 = t2.next
	}

	p3 := PolyAdd(p1, p2)
	p4 := PolyMutiply(p1, p2)

	p4.FormatOutPutPoly()
	p3.FormatOutPutPoly()
}
