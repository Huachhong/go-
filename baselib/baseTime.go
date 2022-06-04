package baselib

import (
	"fmt"
	time2 "time"
)

func TestTime() {
	fmt.Println("########测试常见包time")
	time := time2.Now()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second())
	timestamp := time.Unix()
	fmt.Printf("current timestamp: %d\n", timestamp)

	dateTime := DateTime(1651917520)
	fmt.Printf("timestamp 转化为时间格式: %s\n", dateTime)
	fmt.Println("now time format, time:" + time.Format("2006-01-02 15:04:05"))
}

func DateTime(timestamp int64) string {
	time := time2.Unix(timestamp, 0)
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second())
}
