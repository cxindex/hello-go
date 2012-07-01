package main

//import "fmt"
//and again, we can use slices here
func maxint(a ...int) (tmp int) {
	tmp := 0
	for _, v := range a {
		if v > tmp {
			tmp = v
		}
	}
	return
}

func minint(a ...int) (tmp int) {
	tmp := 0
	for _, v := range a {
		if v < tmp {
			tmp = v
		}
	}
	return
}

func main() {
	println(maxint(1, 2, 3, -4, 999, 0))
	println(minint(1, 2, 3, -4, -999, 0, -100500))
}
