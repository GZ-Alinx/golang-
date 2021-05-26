package main

import (
	"bufio"
	"fmt"
	"os"
)

// 实时输出 标准输入的数据
func Scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	//var sum int
	for scanner.Scan() {
			fmt.Println(scanner.Text())

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}



//func main() {
//	Scanner()
//}