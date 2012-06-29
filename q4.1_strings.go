package main

import "fmt"

func main() {
	s := "A"
	for i := 0; i < 10; i++ {
		fmt.Println(s)
		s = s + "A"
	}
}
