package main

import "fmt"


func main() {
//	var f = make([]float64, 10) //zeroed
//	var f = [...]float64{1,2,3}
//	f := [...]float64{1,2,3} //or
//	f := make([]float64, 10)
	
//	f := [...]float64{1,8,3,5,6,7,8,9,10} //or
	f := make([]float64, 0)
	var sum float64
	for i := 0; i < cap(f); i++ {
		sum += f[i]
	}
	sum/=float64(cap(f))
	fmt.Printf("%T %v\n", f, cap(f))
	fmt.Printf("%v\n", sum)
}
