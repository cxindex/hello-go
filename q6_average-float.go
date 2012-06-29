//with a func
package main

import "fmt"

func fcalc(f []float64) float64 {
	var sum float64
	for _, v1 := range f {
		sum += v1
	}
	return sum / float64(len(f))
}

func main() {
	f := []float64{1.02, 7.3, 9.9}
//	f := make([]float64, 10)
	fmt.Println(fcalc(f))
}
