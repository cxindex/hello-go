package main
//nya nya nya
//see the difference
import "fmt"

func mult(f []interface{}) interface{} {
	fmt.Printf("START TYPE: %T\n", f)
	for i, v := range f {
		fmt.Printf("in TYPE: %T\n", v)
		switch v.(type) {
		case int:
			fmt.Println("OH HAI")
			f[i] = 1 //wont work
		case []int:
			fmt.Println(v.([]int))
		case *[]int:
			fmt.Println(v.(*[]int))
		case []string:
			fmt.Println(len(v.([]string)))
		}
	}
	return f
}

func main() {
	s := []interface{}{"hello", "my", "world"}
	mult(s)
	
	a := []interface{}{1,2,3,43,4}
	mult(a)
	fmt.Println(a)
}