package main
import "fmt"


func main(){
	var a = [10]byte{1,2,3}
	fmt.Println(cap(a), len(a))

}