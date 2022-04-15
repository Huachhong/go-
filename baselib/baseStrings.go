package baselib

import (
	"fmt"
	"strings"
)

/**
常见库 strings
*/

func TestStrings() {
	fmt.Println("########测试常见包strings")
	str := "hello word, 早上好，special this is who am i,晚安，世界"
	fmt.Println(strings.Count(str, "o"))
	fmt.Println(strings.Contains(str, "word"))
	fmt.Printf("ContainsAny, %v\n", strings.ContainsAny(str, "很晚了"))

	fmt.Printf("split,%#v\n", strings.Split(str, "s"))

	f1, f2 := 210.453, 3432322690.5768235
	fmt.Printf("f1=%f, f2=%f\n", f1, f2)

	fmt.Printf("f1=%f, f2=%9f\n", f1, f2)
}
