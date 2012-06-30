//with a func
package main

import "fmt"

func srt(x, y int) (int, int) {
	if x < y {
		return x, y
	}
	return y, x
}

func main() {
	fmt.Println(srt(40, 31))
}
