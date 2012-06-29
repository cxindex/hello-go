package main


import (
	"unicode/utf8"
	"fmt"
)
func main() {
	s := "asdwq11e.... ASDzxcqwe dsk"
	b := []byte(s)
	r := []rune(s)

	for i := 0; i < len(b); i++ {
		switch i {
		case 4:
			r[i] = 'A'
		case 5:
			r[i] = 'B'
		case 6:
			r[i] = 'C'
		}
		// if i >= 4  && i <= 6 {
		// 	r[i] = 'z'
		// }
		
	}
	fmt.Printf("%v %v %v\n", len(s), len(b), utf8.RuneCount(b))
	fmt.Println(string(r))
}
