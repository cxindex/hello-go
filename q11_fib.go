//something with variadic parametrs
package main

import "fmt"

//btw we can do the same thing with slice.append
func fib(x int64) []int64 {
	a := make([]int64, x)
	a[0], a[1] = 1, 1
	for i := 2; int64(i) < x; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	return a
}

func main() {
	a := fib(10)
	fmt.Println(a)
}
