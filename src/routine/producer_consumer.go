package main

import (
	"strconv"
	"sync"
	"time"
)

func producer(cha chan string) {
	time.Sleep(time.Duration(10) * time.Second)
	for i := 0; i < 10; i++ {
		str := "no " + strconv.Itoa(i)
		cha <- str
	}
}

func consumer(cha chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		str := <-cha
		println(str)
	}
}

// 以channel模拟生产者消费者模式
func main() {
	var wg sync.WaitGroup

	cha := make(chan string)
	go producer(cha)
	wg.Add(1)
	go consumer(cha, &wg)

	wg.Wait()
}
