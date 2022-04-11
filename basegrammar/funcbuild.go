package basegrammar

import "fmt"

//函数测试 20220404




func AutoFun(fn func() int) int {
	return fn()
}

func testSlice(s []int) int {
	s[0] = 100
	return 10
}

func TestFun() {
	a := AutoFun(func() int {return 100 })
	fmt.Printf("\na=%d", a)

	ar := [5]int{1, 2, 3, 4, 5}
	s1 := ar[1:4]
	fmt.Printf("\ns1=%v", s1)
	testSlice(s1)
	fmt.Printf("\ns1=%v", s1)
	fmt.Printf("\nar=%v", ar)

	//defer
	var da [5]struct{}
	for i := range da {
		defer fmt.Printf("\ni=%d", i)
	}

}

