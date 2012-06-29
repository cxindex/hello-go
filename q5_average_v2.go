package main

import "fmt"


func main() {
//	f := make([]float64, 0)
	f := [...]float64{1,8,3,5,6,7,8,9,10} //or
	var sum float64
	for v, v2 := range f {
		fmt.Println(v, v2)
		sum += v2
	}
	sum/=float64(cap(f))
	fmt.Printf("%T %v\n", f, cap(f))
	fmt.Printf("%v\n", sum)
}
