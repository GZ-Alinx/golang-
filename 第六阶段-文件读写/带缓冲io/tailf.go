package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("usage: tailf web.log")
	}

	file,err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}



	// 另外一种方式
	reader := bufio.NewReader(file)
	reader := bufio.NewReaderSize(file, 16)


	line := make([]byte, 4096)
	for {
		data, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Println(err)
				break
			}-
			// 程序等待n秒
			time.Sleep(time.Second)
		}else if isPrefix {
			fmt.Println(string(data), isPrefix)
			line = append(line, data...)
		} else {
			if len(line) > 0 {
				fmt.Println(string(append(line,data...)))
				line = make([]byte, 0,4096)
			}else {
				fmt.Println(string(data))
			}
		}
	}
}
