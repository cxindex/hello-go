package main

import "fmt"

func main() {
	p := false
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("Fizz")
			p = true
		}
		if i%5 == 0 {
			fmt.Printf("Buzz")
			p = true
		}
		if !p {
			fmt.Printf("%v", i)
		}
		p = false
		fmt.Println()


	}
}
