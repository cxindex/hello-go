package main

import ( "os"; "fmt"; "bufio"; "time" )

func p() {
	fmt.Println("HI")
}

func rd(s string, reader *bufio.Reader) {
	s, err := reader.ReadString('\n')
	if err == nil {
		fmt.Printf("ASDASD %v", s)
	}
}

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var str string
	for {
		go rd(str, reader)
		p()
		time.Sleep(1000 * time.Millisecond)
	}
}
		

