package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file")
	if err != nil {
		fmt.Println(err)
	}

	data := make([]byte, 1024)
	var next int64
	for {
		count, err := file.ReadAt(data, int64(next))
		if err != nil {
			println()
			break
		}
		next += int64(count)
		//		fmt.Printf("read %d bytes: %q\n", count, data[:count])	
		fmt.Printf("%s", data[:count])
	}
	//	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
