package main

import (
	"fmt"
	"time"
)

/*
func main() {
	os := "mac"
	switch os {
	case "mac":
		fmt.Println("mac!!")
	case "windows":
		fmt.Println("windows!!")
	default:
		fmt.Println("windows!!")
	}
}
*/

/*
func getOs() string {
	return "mac"
}

func main() {
	os := getOs()
	switch os {
	case "mac":
		fmt.Println("mac!!")
	case "windows":
		fmt.Println("windows!!")
	default:
		fmt.Println("windows!!")
	}
}
*/

/*
func getOs() string {
	return "mac"
}

func main() {
	switch os := getOs(); os { //osをswitch文内のみで使用する場合
	case "mac":
		fmt.Println("mac!!")
	case "windows":
		fmt.Println("windows!!")
	default:
		fmt.Println("windows!!")
	}
}
*/

func main() {
	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() >= 12:
		fmt.Println("evening")
	}
}
