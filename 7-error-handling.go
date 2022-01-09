package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("1-if.go")
	if err != nil {
		log.Fatalln("error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("error")
	}
	fmt.Println(count, string(data))
}
