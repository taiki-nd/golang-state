package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {

	result := by2(11)
	if result == "ok" {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}

	if result2 := by2(20); result2 == "ok" { //１行で書く方法
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}

	/*
		num := 9
		if num%2 == 0 {
			fmt.Println("by 2")
		} else if num%3 == 0 {
			fmt.Println("by 3")
		} else {
			fmt.Println("else")
		}
	*/
	x, y := 12, 11
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}
	if x == 10 || y == 10 {
		fmt.Println("OR")
	}
}
