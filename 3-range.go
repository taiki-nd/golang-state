package main

import "fmt"

func main() {
	l := []string{"python", "ruby", "java"}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}
	for i, v := range l {
		fmt.Println(i, v)
	}

	for _, v := range l { //index番号なしで中身だけ取り出し記述
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "orange": 200}

	for k, v := range m {
		fmt.Println(k, v) //key and value
	}

	for v := range m {
		fmt.Println(v) //sliceとmapでは書き方が異なる（_がいらない）
	}

	for _, v := range m {
		fmt.Println(v) //注意
	}
}
