package main

import (
	"fmt"
)

func main() {
	params := &IntOarams{
		P1: 2,
		P2: 3,
	}
	fmt.Println(params.Add())
}
