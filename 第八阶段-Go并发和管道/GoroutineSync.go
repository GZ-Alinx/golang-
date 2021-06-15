package main


import (
	"fmt"
	"sync"
)

/* sync包 提供的其中并发工具类

并发工具类		说明
Mutex			互斥锁
RWMutex			读写锁
WaitGroup		并发等待组
Map				并发安全字典
Cond			同步等待条件
Once			只执行一次
Pool			临时对象池

*/

// 指针类型进程等待

// 计数器 计算
func tests(wg *sync.WaitGroup) {
	// 等待函数执行结束后wg.Done() 计数器 -1
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

func main() {
	// 定义计数器
	var wg sync.WaitGroup
	wg.Add(	1) // 添加计数器 +1
	go tests(&wg) // 传指针




	wg.Wait() // 等待计数器归 0 然年结束进程
}



