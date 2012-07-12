//based on q22.1_cat* both
//-n IS here
package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"bufio"
)

var nflag *bool = flag.Bool("n", false, "number all output lines")

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
			file.Close()
			break
		}
		next += int64(count)
		fmt.Printf("%s", data[:count])
	}
	println()
}

func ncat(s string) {
	fmt.Println("NCAT:", s)
	file, err := os.Open(s)
	if err != nil {
		fmt.Println(err)
	}
	r := bufio.NewReader(file)
	i := 1
	for {
		data, err := r.ReadBytes('\n')
		fmt.Printf("%v   %s", i, string(data))
		i++
		if err == io.EOF {
			file.Close()
			break
		}
	}
	println()
}

func main() {
	flag.Parse()
	for i := 0; i < flag.NArg(); i++ {
		if *nflag {
			ncat(flag.Arg(i))
		} else {
			fcat(flag.Arg(i))
		}
	}

}
