package main

import "fmt"

func worker(ch chan struct{}) {
	<-ch
	fmt.Println("do something")
	close(ch)
}

func main() {
	fmt.Println("do something")
	ch := make(chan struct{})
	go worker(ch)
	ch <- struct{}{}
}