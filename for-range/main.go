package main

import ("fmt")


func main()  {
	words := []string{"Go", "语言", "高性能", "编程"}
for i, s := range words {
    words = append(words, "test")
    fmt.Println(i, s)
}

m := map[string]int{
    "one":   1,
    "two":   2,
    "three": 3,
}
for k, v := range m {
    delete(m, "two")
    m["four"] = 4
    fmt.Printf("%v: %v\n", k, v)
}

ch := make(chan string)
go func() {
    ch <- "Go"
    ch <- "语言"
    ch <- "高性能"
    ch <- "编程"
    close(ch)
}()

for n := range ch {
    fmt.Println(n)
}
}