package basegrammar

import "fmt"

//测试流程语句
func FlowTest() {
	//if
	var a uint8
	a = 30
	if fmt.Printf("a = %d", a); a < 20 {
		fmt.Printf("\n%d < 20", a)
	} else {
		fmt.Printf("\n%d > 20", a)
	}
	//switch
	var b = 100
	fmt.Println()
	switch b {
	case 1:
		fmt.Println("我是1")
	case 2, 3:
		fmt.Println("我是2")
	default:
		fmt.Println("我是default")
	}
	//验证select
	//for
	//range
	s := "abc"
	for i, val := range s {
		fmt.Printf("\n index = %d, value = %c", i, val)
	}
}