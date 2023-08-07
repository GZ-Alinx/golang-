package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/robfig/cron"
)

var (
	Params  []string
	Command string
)

func main() {
	// 创建一个新的Cron调度器
	c := cron.New()

	// 添加定时任务
	// 每秒执行一次
	c.AddFunc("* * * * * *", func() {
		fmt.Println("Step 1 Running cron job at:", time.Now())
		Start_process()
	})

	// 开始Cron调度器
	c.Start()

	// 关闭Cron调度器（优雅地停止）
	defer c.Stop()

	// 主程序可以继续执行其他任务
	fmt.Println("start peocess...")
	time.Sleep(20 * time.Second)
}

// execute command func
func Command_Run(command string, params []string) {
	// 创建Cmd结构体
	cmd := exec.Command(command, params...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("result:  ", string(output))
}

func Start_process() {
	Params = []string{"-lhra"}
	Command = "ls"
	Command_Run(Command, Params)
}

func Stop_process() {
	Params = []string{"-lhra"}
	Command = "ls"
	Command_Run(Command, Params)
}
