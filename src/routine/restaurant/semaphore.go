package main

import (
	"fmt"
	"sync"
	"time"
)

func user(semaphore *sync.WaitGroup) {
	semaphore.Add(1)

	println("im here")
}

// 使用 channel 实现信号量 semaphore
func main() {
	var wg sync.WaitGroup

	// 建立一个容量为2的channel 最多允许2个并发同时执行
	sem := make(chan int32, 2)
	taskNum := 10
	for i := 0; i < taskNum; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			// 向channel写入数据 获取信号
			sem <- 0
			// 从channel取出数据 释放信号
			defer func() { <- sem }()

			// do something for task
			time.Sleep(time.Duration(2) * time.Second)
			fmt.Println(id, time.Now())
		}(i)
	}
	wg.Wait()
}
