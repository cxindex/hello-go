package main

import "fmt"

func main() {
	s := "hello, привет 世界"
	bs := []byte(s)
	fmt.Println(len(s), len(bs))
}