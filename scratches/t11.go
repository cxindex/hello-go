package main
import "fmt"

func slmod(sl []int) {
	for i, _ := range sl {
		sl[i] = 1
	}
}

func main() {
	sl := []int{1,2,3}
	sl1 := make([]int, 10)
	fmt.Printf("%T %v\n", sl, sl)
	fmt.Printf("%T %v\n", sl1, sl1)
	slmod(sl)
	fmt.Println(sl)
}