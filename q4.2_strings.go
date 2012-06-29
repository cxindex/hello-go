package main


import (
	"unicode/utf8"
	"fmt"
)
func main() {
	s := "asdwq11e.... ASDzxcqwe dsk"
	b := []byte(s)
	
	for i := 0; i < len(b); i++ {
		fmt.Printf("%T %s %v\n", b[i], string(b[i]), i)

	}
	fmt.Printf("%v %v %v\n", len(s), len(b), utf8.RuneCount(b))
}
