package main
import "fmt"

func test(i int) { fmt.Printf("%T %v", i, i) }

func main(){
	s := "hello"
	test(len(s))
}