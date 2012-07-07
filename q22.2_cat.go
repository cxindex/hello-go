//based on q22.1_cat_v2
//-n isnt here. just too lazy
package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	flag.Parse()	
	for i := 0; i < flag.NArg(); i++ {
		fmt.Println("FLAG:", flag.Arg(i))
		
		file, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Println(err)
		}

		fi, _ := os.Stat(flag.Arg(i))
		data := make([]byte, fi.Size())
		var next int64
		for {
			count, err := file.ReadAt(data, int64(next))
			if err != nil {
				println()
				break
			}
			next += int64(count)
			fmt.Printf("%s", data[:count])
		}
	}

}
