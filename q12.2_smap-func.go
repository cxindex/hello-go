//time to use some pointers!
package main

import "fmt"

func smapmod(f func(*string), s []string) []string {
	for i, _ := range s {
		f(&s[i])
	}
	return s
}

func main() {
	f := func (s *string) {
		*s = *s + string('.')
	}	
	s := []string{"hello", "world", "BURN, MOTHERFUCKER, BURN!11"}
	fmt.Println(s)
	fmt.Println(smapmod(f, s))
}