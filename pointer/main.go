package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func main() {
	data := []field{{"one"}, {"two"}, {"three"}}

	for _, v := range data {
		fmt.Println(&v.name)
		go v.print()
	}
	time.Sleep(1 * time.Second)
}
