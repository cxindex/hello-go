package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parse(s *string, c, w *int) {
	*c = *c + len(*s) - 1
	*w = *w + len(strings.Fields(*s))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var c, w, l int
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Chars: %v Words: %v Lines: %v\n", c, w, l)
			return
		}
		parse(&s, &c, &w)
		l++
	}
}
