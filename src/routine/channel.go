package main

import "fmt"

// 一般channel的声明形式为：var channelName chan ElementType
// 在channel的用法中，最常见的包括写入和读出。
// 		将一个数据写入至channel的语法: ch <- value
//		从channel中读取数据的语法: value := <- ch
//
// 		向channel写入数据通常会导致程序阻塞，直到有其他 goroutine 从这个channel中取出数据。
// 		如果channel之前没有写入数据， 那么从channel中取出数据也会导致程序阻塞，直到channel中被写入数据为止。

func count(ch chan int) {
	fmt.Println("Counting")
	// 写入时阻塞 直到有其他goroutine取走数据
	ch <- 1
	ch <- 2
}

func main() {
	// 5 个channel的数组
	chs := make([]chan int, 5)

	for i := 0; i < 5; i++ {
		// 将每个channel分配给count函数
		chs[i] = make(chan int)
		go count(chs[i])
	}

	for _, ch := range (chs) {
		// 取出时阻塞 直到其他goroutine写入数据
		value := <- ch
		println(value)
		value2 := <- ch
		println(value2)
	}

}
