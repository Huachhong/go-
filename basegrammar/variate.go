package basegrammar

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Myvariate() {
	fmt.Println("hello")
}

func VariableAndConst() {
	fmt.Println("变量............")
	//普通变量声明
	var v1 int = 1
	var v2 int
	v3 := 3 //短标签声明
	//批量声明
	var (
		v4 = 5
		v5 = 6
	)
	fmt.Printf("\nv1=%v,v2=%v,v3=%v,v4=%v,v5=%v", v1,v2,v3,v4,v5)
	fmt.Println("\n常量............")
	const cn1  = 1
	const cn2  = "con string"
	//iota是go语言的常量计数器，只能在常量的表达式中使用。iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用
	const (
		cn3 = iota
		cn4
		cn5
	)
	fmt.Printf("\ncn1=%v,cn2=%v,cn3=%v,cn4=%v,cn5=%v",cn1, cn2,cn3,cn4,cn5)
	fmt.Println("\n浮点型............")
	var f1 = 1.0
	var f2 float32 = 1.02
	fmt.Printf("f1=%f,f2=%f",f1, f2)

	fmt.Println("\n字符类型............")
	//字符用单引号表示
	var c1 byte = 1
	var c2 rune = 'a'
	var c3 rune = '中'
	fmt.Printf("\nc1=%v, type is %T", c1, c1)
	fmt.Printf("\nc2=%v, type is %T", c2, c2)
	fmt.Printf("\nc3=%v, type is %T", c3, c3)
	//大小写转换
	c4 := 'A' - 'a'
	c5 := 'g'
	c6 := c5 + c4
	fmt.Printf("c6= %v,c6=%c,type is %T", c6, c6, c6)

	fmt.Println("\n强制类型转换............")
	v1 = int(f2)
	fmt.Printf("\nv1=%v,type is %T", v1, v1)

	fmt.Println("\n布尔类型............")
	var b1 bool
	var b2 bool = true
	fmt.Printf("\nb1=%v, type is %T", b1, b1)
	fmt.Printf("\nb2=%v, type is %T", b2, b2)

	fmt.Println("\n字符串类型............")
}

//计算及遍历字符串长度
func RandString() {
	s1 := "hello,中国！"
	fmt.Printf("len=%d,utf8 len=%d,%d,%d\n", len(s1), strings.Count(s1, "") - 1, len([]rune(s1)),utf8.RuneCountInString(s1))
	for _, r := range s1 {
		fmt.Printf("%v(%c),",r,r)
	}
}

//数组
func ArrTest() {
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr2 = [5]int{1, 2, 3, 4, 5}
	var arr3 = [...]int{1, 2, 3}
	var arr4 = [5]string{2:"where", 4:"are"}
	arr5 := [3]string{0:"yh"}
	arr6 := [...]struct{
		name string
		age uint8
	}{{"小明", 10}, {"小强",9}}
	fmt.Println(arr1, arr2, arr3, arr4, arr5, arr6)
	fmt.Println("遍历数组")
	for  _, v := range arr1 {
		fmt.Printf("val = %v\n", v)
	}
}

func SliceTest() {
	fmt.Println("切片...")
	//声明切片
	var s1 []int
	fmt.Printf("s1=%v", s1)
	fmt.Println("切片...make创建切片")
	var s2 []int = make([]int, 2, 10)
	fmt.Printf("s2=%v", s2)
	s3 := [5]int{1, 2, 3, 4, 5}
	s4 := s3[1:3]
	s5 := s3[:3]
	fmt.Printf("\nlen=%d, cap=%d, s4=%v,s5_cap=%d, s5=%v", len(s4), cap(s4), s4, cap(s5), s5)
	s4[0] += 2
	s6 := append(s5, 1)
	fmt.Printf("\ns3=%v, s4=%v, s6=%v", s3, s4, s6)
}

func MapTest() {
	var m1 = make(map[string]string, 2)
	m1["name"] = "李小飞"
	fmt.Printf("m1=%v, m1 type=%T", m1, m1)
	//在声明的时候初始化
	m2 := map[string]string{
		"na":"明",
		"age":"20",
	}
	fmt.Printf("\nm2=%v", m2)
	//判断某个键是否存在。// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	val, ok := m2["na"]
	if ok {
		fmt.Printf("\nval=%v", val)
	} else {
		fmt.Printf("\n键值不存在")
	}
	//遍历
	for key, val := range m2 {
		fmt.Printf("\nkey = %v, val = %v", key, val)
	}
	//使用delete()内建函数从map中删除一组键值对 格式；delete(map, key)

	//元素为map类型的切片
	var sliceMap = make([]map[string]string, 2)
	sliceMap[0] = make(map[string]string, 10) //需要先进行初始化
	sliceMap[0]["name"] = "hello word"
	sliceMap[0]["addr"] = "广州"
	sliceMap[1] = make(map[string]string, 10) //需要先进行初始化
	sliceMap[1]["name"] = "hello go"
	sliceMap[1]["addr"] = "深圳"
	fmt.Printf("\nsliceMap=%v", sliceMap)

	//值为切片的map
	var mapSlice = make(map[string][]string, 10)
	fmt.Printf("\nmapSlice=%v", mapSlice)
	mapSlice["goods"] = []string{"苹果","栗子"}
	mapSlice["prices"] = []string{"10", "20"}
	fmt.Printf("\nmapSlice=%v", mapSlice)
}
