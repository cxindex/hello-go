package main
import "fmt"

func main(){
	var v int64 = 0x80000000
	var i uint64
	for ; i<=5; i++ {
		fmt.Printf("%v %x\n", v, v>>i)
//		fmt.Println(i)
	}
}
	