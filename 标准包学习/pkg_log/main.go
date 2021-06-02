package main

import (
	"log"
)

func test() {
	defer func() {
		recover()
	}()
	log.Panicln("panic")
}

func testFatal() {
	log.Fatalln("fatalln ")
}

func main() {
	log.SetPrefix("日志前缀: ") // 设置日志前缀
	log.Print("00 我是Alinx")
	log.Println("01 我是Alinx")
	log.Printf("日志：%s", "02 我是alinx")

	// 日志格式 Ldate
	log.SetFlags(log.Ldate)
	log.SetFlags(log.Ltime)
	test()
	testFatal()
	log.Println("over")


}
