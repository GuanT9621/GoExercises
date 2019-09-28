package main

import (
	"sync"
	"time"
)

func add(count *int32) {
	*count += 1
}

func safeAdd(count *int32, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	*count += 1
}

// 启动两个协程 分别给 count 加 1 50000次 结果可见是线程不安全的
// 上了 Mutex 锁 就实现了线程安全了
func main() {
	count := int32(0)
	safeCount := int32(0)
	var mutex sync.Mutex

	loop := 50000
	for i := 0; i < loop; i++ {
		go add(&count)
		go add(&count)
		go safeAdd(&safeCount, &mutex)
		go safeAdd(&safeCount, &mutex)
	}

	time.Sleep(time.Duration(5) * time.Second)

	println("unsafeCount ", count)
	println("safeCount ", safeCount)
}
