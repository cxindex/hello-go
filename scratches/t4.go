package main

import "fmt"

func a(arg ...int) []int { return arg }

func main() {
	ar := []int{1,3,4}
	fmt.Println(a(ar...))

	av := 1
	fmt.Println(a(av))
}
