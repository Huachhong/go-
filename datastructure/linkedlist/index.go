package linkedlist

import "fmt"

func Test() {
	fmt.Printf("%v", new(Node))
	fmt.Println(head)
	head = nil
	traverse(head)
	addNode(head, -1)
	addNode(head, 0)
	addNode(head, 10)
	addNode(head, 5)
	addNode(head, 8)
	traverse(head)

	fmt.Println(lookupNode(head, 10))
	Delete(head, 10)
	traverse(head)

	//双向链表
	fmt.Println("-> 双向链表!")
	fmt.Println(dhead)
	dhead = nil
	addDNode(dhead, 10)
	addDNode(dhead, 30)
	addDNode(dhead, 8)
	addDNode(dhead, 20)
	addDNode(dhead, 6)
	addDNode(dhead, 12)
	addDNode(dhead, 15)
	dhead.dtraverse(dhead)
	//fmt.Println()
	fmt.Println(dsize(dhead))
}
