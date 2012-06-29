package main

import "fmt"

func main() {
	var a [10]int
	for i := 0; i < len(a); i++ {
		a[i] = i
		fmt.Println(a[i], "I:", i)
	}
	fmt.Println(a)
}
