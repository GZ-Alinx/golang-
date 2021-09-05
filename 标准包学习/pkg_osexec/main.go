package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 查找指令绝对路径
	cmd, err := exec.LookPath("ping")
	if err != nil {
		fmt.Println(cmd)
	}

	// 默认是阻塞的,可以加start 和Run 变成异步执行
	// Wait 是等待执行结束
	c := exec.Command(cmd, "www.baidu.com")
	ret, err := c.Output()
	c.Wait()
	fmt.Println(string(ret), err)
}
