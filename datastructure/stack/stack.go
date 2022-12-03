package stack

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var size = 0
var stack = new(Node)

// Push 进栈
func Push(v int) bool {
	//	空栈
	if stack == nil {
		stack = &Node{v, nil}
		size = 1
		return true
	}
	// 	否则将插入节点作为栈的头结点
	temp := &Node{v, nil}
	temp.Next = stack
	stack = temp
	size++
	return true
}

func Pop(t *Node) (int, bool) {
	//	空栈
	if size == 0 {
		return 0, false
	}
	// 只有一个节点
	if size == 1 {
		size = 0
		stack = nil
		return t.Value, true
	}
	//	将栈的头结点指针指向下一个节点，返返回当前的节点数
	stack = stack.Next
	size--
	return t.Value, true
}

//	循环 （不删除任何节点，只读取值)
func traverse(t *Node) {
	if size == 0 {
		fmt.Println(" -> 空栈")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}
