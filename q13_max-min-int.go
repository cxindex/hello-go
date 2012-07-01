package main

//import "fmt"
//and again, we can use slices here
func maxint(a ...int) int {
	tmp := 0
	for _, v := range a {
		if v > tmp {
			tmp = v
		}
	}
	return tmp
}

func minint(a ...int) int {
	tmp := 0
	for _, v := range a {
		if v < tmp {
			tmp = v
		}
	}
	return tmp
}

func main() {
	println(maxint(1, 2, 3, -4, 999, 0 ))
	println(minint(1, 2, 3, -4, -999, 0, -100500))
}