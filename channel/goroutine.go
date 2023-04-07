package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("-- new goroutine task: i = %d\n", i)
		time.Sleep(3 * time.Second) //延时1s
		if i >= 20 {
			break
		}
	}
}

/**
 * 主任务和 go func(){} 应用
 */
func main() {
	go newTask() //新建一个协程任务
	i := 0
	for {
		i++
		fmt.Printf("main task: i = %d\n", i)
		time.Sleep(2 * time.Second) //延时1s
		if i >= 10 {
			break
		}
	}
	//这里被for阻塞
	fmt.Printf("main OK!\n")
}
