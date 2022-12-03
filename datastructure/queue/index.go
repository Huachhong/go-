package queue

import "fmt"

func Test() {
	//链式队列
	fmt.Println("链式队列")
	linkqueue := new(LinkQueue)
	linkqueue.Push(1)
	linkqueue.Push(8)
	linkqueue.Push(9)
	linkqueue.Push(5)
	linkqueue.Push(15)
	linkqueue.Push(2)
	linkqueue.Push(3)
	linkqueue.traverse()
	fmt.Println()
	v1, _ := linkqueue.Pop()
	fmt.Println("pop -> ", v1, " size:", linkqueue.size)
	v2, _ := linkqueue.Pop()
	fmt.Println("pop -> ", v2, " size:", linkqueue.size)
	linkqueue.traverse()

}
