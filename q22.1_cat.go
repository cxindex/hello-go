package main

import ("fmt"; "os"; "log"; "bufio"; "io")

func main() {
	file, err := os.Open("file") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	for {
		buf, err := reader.ReadBytes('\n')
		fmt.Printf("%s", string(buf))
		if err == io.EOF {
			fmt.Println()
			break			
		}
	}
	
}