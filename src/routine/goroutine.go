package main

import "time"

/**
https://studygolang.com/articles/11627
在golang中均为协程，包括 main 和所有以 go 命令启动的函数。
*/

func sum(number1, number2 int32) int32 {
	println("in sum")
	return 0
}

func main() {
	println("start main")
	go sum(1,2)
	time.Sleep(time.Duration(1) * time.Second)
	println("end main")
}