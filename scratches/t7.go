package main

import "fmt"

func ta(a []int64) {
	fmt.Printf("%T, %v\n", a, 1)
}

func tap(a *[]int64) {
	fmt.Printf("%T, %v\n", a, 1)
}


func main() {
	a := make([]int64, 999999999)
	for i, _ := range a {
		a[i] = int64(i)
	}
	
	ta(a)
}