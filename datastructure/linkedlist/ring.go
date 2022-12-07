package linkedlist

import "fmt"

//双向循环链表

type Ring struct {
	Value      interface{} //数据
	next, prev *Ring       //	前驱和后驱节点
}

//初始化空的循环链表,前驱和后驱都指向自己
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

//	创建N个节点的循环链表
func New(n int) *Ring {
	if n < 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}
func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

// 	获取第n个节点
//	因为链表是循环的，当 n 为负数，表示从前面往前遍历，否则往后面遍历：
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

//	往节点A，链接一个节点，并且返回之前节点A的后驱节点
func (r *Ring) Link(s *Ring) *Ring {
	/*	n := r.Next()
		if s != nil {
			p := s.Prev()
			r.next = s
			s.prev = r
			n.prev = p
			p.next = n
		}
		return n*/
	n := r.Next()
	if s != nil {
		r.next = s
		s.prev = r
		s.next = n
		n.prev = s
	}
	return n
}

func (r *Ring) traverse() {
	node := r
	for {
		fmt.Printf("%d -> ", node.Value)
		node = node.Next()
		if node == r {
			fmt.Printf("\nnode:%v,r:%v", node, r)
			break
		}
	}
	fmt.Println()
	/*
		for {
			fmt.Printf("pre:%v, %d -> , next:%v\n", node.prev.Value, node.Value, node.next.Value)
			node = node.Next()
			if node == r {
				fmt.Printf("\nnode:%v,r:%v", node, r)
				break
			}
		}*/
}

func RingTest() {
	r := &Ring{Value: -1}
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})

	r.traverse()
}
