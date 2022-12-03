package queue

import (
	"errors"
	"fmt"
	"sync"
)

//	顺序队列(数组实现)

type ArrayQueue struct {
	array []int
	size  int
	lock  sync.Mutex
}

func (queue *ArrayQueue) Push(v int) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	//	放入切片中
	queue.array = append(queue.array, v)
	//	队列元素数量+1
	queue.size++
}

func (queue *ArrayQueue) Pop() (int, error) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if queue.size == 0 {
		fmt.Println("-> 空队列")
		return 0, errors.New("空队列")
	}
	v := queue.array[0]

	newArray := make([]int, queue.size-1, queue.size-1)
	for i := 1; i < queue.size; i++ {
		newArray[i-1] = queue.array[i]
	}
	queue.array = newArray
	queue.size--
	return v, nil
}

//	链式队列(链表实现)

type LinkQueue struct {
	root *LinkNode  //	链表起点
	size int        //	队列元素数
	lock sync.Mutex //	为了并发安全使用的锁
}

type LinkNode struct {
	Value int
	Next  *LinkNode
}

//	入栈

func (queue *LinkQueue) Push(v int) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.root == nil {
		queue.root = new(LinkNode)
		queue.root.Value = v
	} else {

		//	一直遍历到链表尾部
		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}
		//新节点放在链表的尾部
		nowNode.Next = &LinkNode{v, nil}
	}
	queue.size++
}

func (queue *LinkQueue) Pop() (int, error) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if queue.size == 0 || queue.root == nil {
		fmt.Println("链表队列为空")
		return 0, errors.New("链表队列为空")
	}
	//	顶部元素出队
	topNode := queue.root
	v := topNode.Value

	//	将顶部元素的后继链接接上
	queue.root = topNode.Next
	queue.size--
	return v, nil
}

func (queue *LinkQueue) traverse() {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	nowNode := queue.root
	for nowNode.Next != nil {
		fmt.Printf("%d -> ", nowNode.Value)
		nowNode = nowNode.Next
	}
	fmt.Printf("%d -> ", nowNode.Value)
}
