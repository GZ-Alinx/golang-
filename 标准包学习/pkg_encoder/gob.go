package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	users := []*User{
		{1, "1", "sz"},
		{2, "2", "xx"},
	}
	file, _ := os.Create("data.log")

	defer file.Close()

	encoder := gob.NewEncoder(file)
	err := encoder.Encode(users)
	fmt.Println(err)

	file, _ = os.Open("datagob")
	defer file.Close()

	//
	users2 := make([]*users, 0)
}
