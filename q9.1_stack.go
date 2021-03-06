package main

import "fmt"

type stack struct {
	i int
	a [10]int
}

func (st *stack) push(x int) int {
	if st.i < len(st.a) {
		st.a[st.i] = x
		st.i++
		return 0
	} else {
		fmt.Println("Stack is full")
	}
	return 1
}

func (st *stack) pop() int {
	if st.i > 0 {
		st.i--
		st.a[st.i] = 0
	}
	return 0
}

func main() {
	var st stack
	for i := 0; i < 18; i++ {
		st.push(i+10)
	}
	fmt.Println(st.a)
	fmt.Printf("%v\n", st)

	for i := 0; i < 18; i++ {
		st.pop()
		fmt.Println(st.a)
	}
}
