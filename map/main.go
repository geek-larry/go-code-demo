package main

import (
	"fmt"
)

func main() {

	var m = make(map[string]int, 10)
	m["one"] = 1
	for _, v := range m {
		fmt.Println(v)
	}

	data := []int{1, 2, 3}
	for _, v := range data {
		fmt.Println(&v)
		v *= 10
	}
	fmt.Println("data", data)
}
