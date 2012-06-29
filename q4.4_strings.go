package main


import (
	"unicode/utf8"
	"fmt"
)
func main() {
	s := "foobar"
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		fmt.Println(i, j)
		b[i], b[j] = b[j], b[i]
	}
	s = string(b)
	fmt.Printf("%v %v %v\n", len(s), len(b), utf8.RuneCount(b))
	fmt.Println(s)
}
