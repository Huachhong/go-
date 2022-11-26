package basegoroutine

import (
	"fmt"
	"runtime"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine: i=%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func Run() {
	//RunGoroutine()
	RunGoroutineExit()
}

func RunGoroutine() {
	go newTask()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i=%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func RunGoroutineExit() {
	go func() {
		defer fmt.Printf("a.defer\n")
		func() {
			defer fmt.Printf("b.defer\n")
			runtime.Goexit()
			fmt.Printf("B")
		}()
	}()

	for {

	}
}
