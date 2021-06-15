package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"encoding/hex"
)


func main() {
	f, err := os.Open("file.txt")


	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	data := hex.EncodeToString(h.Sum(nil))
	fmt.Println(data)
	//fmt.Printf("%x", h.Sum(nil))
}

