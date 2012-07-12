package main

import (
	"fmt"
	"time"
)

// _, ok) //value, true/false
func more(x, y interface{}) interface{} {
	switch x.(type) {
	case int:
		if _, ok := y.(int); ok {
			fmt.Println(x, y)
		} else {
			panic(fmt.Sprintf("CANT COMPARE %T AND %T", x, y))
		}
	case float64:
		if _, ok := y.(float64); ok {
			fmt.Println(x, y)
		} else {
			panic(fmt.Sprintf("CANT COMPARE %T AND %T", x, y))
		}
	default:
		fmt.Printf("UNKNOWN TYPES %T AND %T\n", x, y)

	}
	return 1
}

func main() {
	more("1.1", "")
	//	more(1.1, 1) //float int - panic
	more(1.1, 1.1)
	more(1, 1)
}}
