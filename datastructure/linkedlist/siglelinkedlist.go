package linkedlist

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var head = new(Node)

func addNode(t *Node, v int) int {
	if head == nil {
		head = &Node{v, nil}
		return 0
	}
	if v == t.Value {
		fmt.Println("节点已经存在")
		return -1
	}
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}
	return addNode(t.Next, v)
}

//	遍历链表
func traverse(t *Node) {
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

func lookupNode(t *Node, v int) bool {
	if t == nil {
		t = &Node{v, nil}
		head = t
		return false
	}
	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}
	return lookupNode(t.Next, v)
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> 空链表!")
	}
	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func Delete(t *Node, v int) bool {

	if t == nil {
		fmt.Println("-> 空链表!")
		return false
	}
	pre := new(Node)
	//next := t
	for t != nil {
		if t.Value == v {
			if pre == nil {
				t = t.Next
			}
			if t.Next == nil {
				pre.Next = nil
			}
			pre.Next = t.Next
			return true
		}
		pre = t
		t = t.Next
	}
	return false
}
