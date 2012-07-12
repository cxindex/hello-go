package main
import "fmt"

func f(t interface{}) {
	switch t.(type) {
		case string:
		fmt.Println("STR")
	}
	
}

func main() {
	str := "hello"

	f(str)
}