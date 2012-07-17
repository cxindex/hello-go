package main

import "fmt"

func main() {
	s := []string{"a", "a", "b", "c", "c", "a", "a", "z"}
	var ps string
	for _, v := range s {
		if v != ps {
			fmt.Printf("%s", v)
		}
		ps = v
	}
	println()
}