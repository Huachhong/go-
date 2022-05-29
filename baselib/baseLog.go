package baselib

import (
	"fmt"
	"log"
)

func TestLog() {
	fmt.Println("########测试常见包Log")
	log.Printf("这是一条普通的日志: %s", "logger")
	log.Fatal("我是Fatal日志")
	fmt.Println("hhhhh")
}
