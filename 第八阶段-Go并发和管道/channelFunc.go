package main

import (
	"bufio"
	"fmt"
	"os"
)

func printInput(ch chan string) {
	// 使用for循环从channel中获取数据
	for val := range ch{
		if val == "EOF"{
			fmt.Println("数据目前为空")
			break
		}
		fmt.Printf("Input is %s\n", val)
	}
}

func main() {
	ch := make(chan string)
	go printInput(ch)


	// 从命令行读取输入
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := scanner.Text()
		ch <- val
		if val == "EOF"{
			fmt.Println("End the game")
			break
		}
	}
	defer close(ch) // channel 关闭后 将不再能往里面写数据 不然会抛出panic异常
	// 而再从关闭的channel中读取数据的阻塞的goroutine将会接收到零值并直接返回

}
