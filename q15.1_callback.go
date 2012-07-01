package main

import "fmt"

func pt() func(int) int {
	return func(x int) int { return x + 2 }
}

func main() {
	f := pt()
	fmt.Printf("%T, %v\n", f, f(2))
}
