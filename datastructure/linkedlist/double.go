package linkedlist

import (
	"fmt"
)

type DNode struct {
	Value int
	Prev  *DNode
	Next  *DNode
}

var dhead = new(DNode)

func addDNode(t *DNode, v int) int {
	if dhead == nil {
		t = &DNode{v, nil, nil}
		dhead = t
		return 0
	}
	if t.Value == v {
		fmt.Println("->节点已经存在")
		return -1
	}
	if t.Next == nil {
		tmp := t
		t.Next = &DNode{v, tmp, nil}
		return -2
	}
	return addDNode(t.Next, v)
}

//	遍历
func (d *DNode) dtraverse(t *DNode) {
	if t == nil {
		fmt.Println("-> 空链表!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func drevtraverse(t *DNode) {
	if t == nil {
		fmt.Println("-> 空链表!")
		return
	}
	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}
	for temp.Prev != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Prev
	}
	fmt.Printf("%d -> ", temp.Value)
	fmt.Println()
}

func dsize(t *DNode) int {
	if t == nil {
		return 0
	}
	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func lookupDNode(t *Node, v int) bool {
	if head == nil {
		return false
	}
	if t.Value == v {
		return true
	}
	return lookupDNode(t.Next, v)
}
