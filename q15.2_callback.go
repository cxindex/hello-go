package main

import "fmt"
//kinda interesting
func ptx(y int) func(int) int {
	return func(x int) int { return x + y }
}

func main() {
	f := ptx(33) //y == 33
	fmt.Printf("%T, %v\n", f, f(9)) //x == 9. i.e. return 33 + 9
}
