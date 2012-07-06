package main

//nya nya nya
//answers says we should do it with a callback, but fuck it
import "fmt"

func mult(f interface{}) interface{} {
	fmt.Printf("%T\n", f)
	switch f.(type) {
	case int:
	// f.(int) = 1 //wont work
	case []int:
		fmt.Println(f.([]int))
	case *[]int:
		fmt.Println(f.(*[]int))
	case []string:
		fmt.Println(len(f.([]string)))
	}
	return f
}

func main() {
	s := []string{"hello", "my", "world"}
	mult(s)

	a := []int{1, 2, 3, 43, 4}
	mult(a)
	mult(&a)
}
