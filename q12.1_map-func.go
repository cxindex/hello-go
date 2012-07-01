//time to use some pointers!
package main

import "fmt"

func mapmod(f func(*int), a []int) []int {
	for i, _ := range a {
		f(&a[i])
	}
	return a
}

func main() {
	f := func (x *int) {
		*x = *x**x
	}	
	a := []int{999, 42, 13}
	fmt.Println(mapmod(f, a))
	fmt.Println(a)
}