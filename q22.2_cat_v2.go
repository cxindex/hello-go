//based on q22.1_cat* both
//-n IS here
package main

import (
	"flag"
	"fmt"
	"os"
)

func fcat(s string) {
	fmt.Println("FCAT:", s)
	file, err := os.Open(s)
	if err != nil {
		fmt.Println(err)
	}
	fi, _ := os.Stat(s)
	data := make([]byte, fi.Size())
	var next int64
	for {
		count, err := file.ReadAt(data, int64(next))
		if err != nil {
			println()
			file.Close()
			break
		}
		next += int64(count)
		fmt.Printf("%s", data[:count])
	}
}

func main() {
	flag.Parse()
	for i := 0; i < flag.NArg(); i++ {
		fcat(flag.Arg(i))
	}

}
