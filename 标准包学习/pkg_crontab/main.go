package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	// 创建一个新的Cron调度器
	c := cron.New()

	// 添加定时任务
	// 每秒执行一次
	c.AddFunc("* * * * * *", func() {
		fmt.Println("Step 1 Running cron job at:", time.Now())
	})

	// 每天的12点整执行
	c.AddFunc("* 0 12 * * *", func() {
		fmt.Println("Step 2 Running cron job at:", time.Now())
	})

	// 开始Cron调度器
	c.Start()

	// 程序终止前，可以让调度器持续运行一段时间，或者等待信号终止
	// time.Sleep(5 * time.Minute)

	// 关闭Cron调度器（优雅地停止）
	defer c.Stop()

	// 主程序可以继续执行其他任务
	fmt.Println("Main program continues...")
	time.Sleep(10 * time.Second)
}
