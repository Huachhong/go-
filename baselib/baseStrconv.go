package baselib

import (
	"fmt"
	"strconv"
)

func TestStrconv() {
	fmt.Println("########测试常见包strconv")
	s1 := "01"
	s2, err := strconv.Atoi(s1)

	if err != nil {
		fmt.Println("string to int error")
	} else {
		fmt.Printf("strconv Atoi Type %T, s2 = %#v\n", s2, s2)
	}

	s3 := 1000
	s4 := strconv.Itoa(s3)
	fmt.Printf("strconv Itoa, Type %T, s4 = %#v\n", s4, s4)

	//Parse系列函数
	b1 := "F"
	b2, err := strconv.ParseBool(b1)
	fmt.Printf("strconv parse, Type %T, b2 = %#v, error=%s\n", b2, b2, err)

	i1 := "-2"
	i2, err := strconv.ParseInt(i1, 10, 64)
	fmt.Printf("strconv parse, Type %T, i2 = %#v, error=%s\n", i2, i2, err)

}
