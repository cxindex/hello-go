package main

import "fmt"

func bs(a []int) []int {
	for sw := true; sw == true; {
		sw = false
		for i := 1; i < len(a); i++ {
			if a[i-1] > a[i] {
				a[i-1], a[i] = a[i], a[i-1]
				sw = true
			}
		}
	}
	return a
}

func main() {
	a := []int{42, 5, 7, 1, 9, 41, 12, 23}
	fmt.Println(bs(a))
}
