package stack

import "fmt"

func Test() {
	stack = nil
	v, b := Pop(stack)
	if b {
		fmt.Print(v, "	")
	} else {
		fmt.Println("Pop() 失败")
	}
	Push(100)
	traverse(stack)
	for i := 0; i < 10; i++ {
		Push(i)
	}
	for i := 0; i < 10; i++ {
		v, b := Pop(stack)
		if b {
			fmt.Print(v, " ")
		} else {
			break
		}
	}
	fmt.Println()
	traverse(stack)
}
