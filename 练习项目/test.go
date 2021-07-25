	package main

	import "fmt"

	func main() {
		var v1,v2,v3 chan int
		var c1,c2 int

		select {
			case c1 = <-v1:
				fmt.Println("收到:", c1, "来自:", v1)
			case c2 = <-v2:
				fmt.Println("收到:", c2, "来自:", v2)
			case c3,ok := (<-v3):
				if ok {
					fmt.Println("收到:", c3, "来自:", v3)
				}else{
					fmt.Println("v3 无数据")
				}
		default:
			fmt.Println("无通信数据")
		}
	}