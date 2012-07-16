package main

import (
	// "fmt"
	"os"
)

func main() {
	f, _ := os.Open("../README.md")
	buf := make([]byte, 100)
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			println()
			break
		}
		// fmt.Printf("%s", string(buf[:n]))
		os.Stdout.Write(buf[:n])
	}

}
