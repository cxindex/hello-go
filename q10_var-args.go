//something with variadic parametrs
package main

import "fmt"

func pl(a ...int) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func main() {
	pl(1,2,3,999,42)
}