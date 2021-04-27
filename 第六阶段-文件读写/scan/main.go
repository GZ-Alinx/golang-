package main

// 不带缓冲/带缓冲 的IO读写  语言层
import "fmt"

func main() {
	var num int
	fmt.Println(fmt.Scan(&num))
	fmt.Println(num)

	fmt.Println(fmt.Scan(&num))
	fmt.Println(num)

	fmt.Println(fmt.Scan(&num))
	fmt.Println(num)

	fmt.Println(fmt.Scan(&num))
	fmt.Println(num)

	fmt.Println(fmt.Scan(&num))
	fmt.Println(num)
}
