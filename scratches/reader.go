package main

import ( "os"; "fmt"; "bufio" )

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	for {
		s, err := reader.ReadString('\n')
		if err == nil {
			fmt.Printf("%v", s)
		}
	}
}
		

