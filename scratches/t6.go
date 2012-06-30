package main

import "fmt"

type stack struct {
	i int
	a [10]int
}

func tfs(st *stack) {
	st.i = 999
}

func tf(x *int) int {
	*x = 33
	return 1
}

func main() {
	x := 1
	tf(&x)
	fmt.Println(x)

	var st stack
	tfs(&st)
	fmt.Println(st.i)
}
