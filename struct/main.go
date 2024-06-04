package main

import (
	"fmt"
	"unsafe"
)

type Door struct{}

func (d Door) Open() {
	fmt.Println("Open the door")
}

func (d Door) Close() {
	fmt.Println("Close the door")
}

func main() {
	fmt.Println(unsafe.Sizeof(struct{}{}))
}