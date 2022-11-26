package baselib

import (
	"fmt"
	"time"
)

func TestTime() {
	fmt.Println("########测试常见包time")
	now := time.Now()
	//fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second())
	//获取当前时间戳
	timestamp := now.Unix()
	fmt.Printf("current timestamp: %d\n", timestamp)

	dateTime := DateTime(1651917520)
	fmt.Printf("timestamp 转化为时间格式: %s\n", dateTime)
	fmt.Println("now time format, time:" + now.Format("2006-01-02 15:04:05"))
	tick_time()
}

func DateTime(timestamp int64) string {
	now := time.Unix(timestamp, 0)
	return now.Format("2006-01-02 15:04:05")
	//return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

func tick_time() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Printf("%v\n", i.Format("2006-01-02 15:04:05"))
	}
}
