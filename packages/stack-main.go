package main
//without $GOPATH
import stack "./stack"
import "fmt"
//and we can use st.I
func main() {
	var st stack.Stack
	for i := 0; i < 18; i++ {
		st.Push(i+10)
	}
	fmt.Println(st.A)
	fmt.Printf("%v\n", st.A)

	for i := 0; i < 18; i++ {
		st.Pop()
		fmt.Println(st.A)
	}
}
