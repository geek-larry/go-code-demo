package main

import (
	"fmt"
)

func main() {

	v := 1
	fmt.Println(&v)
	{
		v := 2
		fmt.Println(&v)
		fmt.Println(v)
		{
			v := 3
			fmt.Println(&v)
			fmt.Println(v)
		}
	}
	fmt.Println(&v)
	fmt.Println(v)
}
