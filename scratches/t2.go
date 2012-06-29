package main
import "fmt"


func main(){
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			fmt.Println("first")
			i++
//			continue
		case 1:
			fmt.Println("second")
		}
	}
}