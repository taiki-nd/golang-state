package main

import (
	"fmt"
	"log"
	"os"
)

/*
func main() {
	defer fmt.Println("world")
	fmt.Println("Hello")
}
*/

/*
func foo() {
	defer fmt.Println("world foo")
	fmt.Println("Hello foo")
}

func main() {
	foo()
	defer fmt.Println("world")
	fmt.Println("Hello")
}
*/

/*
func main() {
	fmt.Println("run")
	defer fmt.Println(1) //スタッキングdeferと言って後ろから実行される
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")
}
*/

func main() {
	file, err := os.Open("./1-if.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //開いたファイルは閉じる必要がある
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(data)
	fmt.Println(string(data))
}
