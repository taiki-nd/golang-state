package main

import "fmt"

/*
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
*/

/*
func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue //ここでforに戻るので３の時にfmt.Println(i)が実行されない
		}
		fmt.Println(i)
	}
}
*/

func main() {

	for i := 0; i < 10; i++ {
		if i == 4 {
			fmt.Println("continue")
			continue //ここでforに戻るので３の時にfmt.Println(i)が実行されない
		}
		if i > 5 {
			fmt.Println("break")
			break //forループを抜ける
		}
		fmt.Println(i)
	}

	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)

}
