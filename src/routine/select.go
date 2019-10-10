package main

// 每个 case 必须是一个通信操作，要么是发送要么是接收。
// select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。
// 一个默认的子句应该总是可运行的。

// select 语句的语法：
//		每个 case 都必须是一个通信
//		所有 channel 表达式都会被求值
//		所有被发送的表达式都会被求值
//		如果任意某个通信可以进行，它就执行，其他被忽略。
//		如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
//		否则：
//			如果有 default 子句，则执行该语句。
//			如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
func main() {
	ch := make(chan int32, 1)
	for i := 0; i < 20; i++ {
		select {
		case ch <- 1:
			// 如果成功向ch写入数据，则进行该case处理语句
			println("case1")
		case ch <- 2:
			// 如果成功向ch写入数据，则进行该case处理语句
			println("case2")
		default:
			// 如果上面都没有成功，则进入default处理流程
			println("default")
		}
		number := <-ch
		println("number =", number)
	}
}
