package main

import "fmt"

func foo(a [2]int) {
	a[0] = 200
}

func foo2(a *[2]int) {
	(*a)[0] = 200
}

func foo1(a []int) {
	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8)
	a[0] = 200
}

func foo3(a *[]int) {
	*a = append(*a, 1, 2, 3, 4, 5, 6, 7, 8)
	(*a)[0] = 200
}

func main() {
	a := [2]int{1, 2}
	foo2(&a)
	fmt.Println(a)


	b := []int{1, 2}
	foo1(b)
	fmt.Println(b)

	c := []int{1, 2}
	foo3(&c)
	fmt.Println(c)
}