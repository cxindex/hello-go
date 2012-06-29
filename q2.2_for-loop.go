package main

import "fmt"

func main() {
	var i int
here:

	fmt.Printf("%v\n", i)
	i++
	if i < 10 {
		goto here
	}
	fmt.Printf("LAST: %v\n", i)
}
