package main
// here we wanna test func as param to func
import "fmt"

//func r(arg int) (int, int) { return arg, arg+4}
func f(f2 func(int) (int, int), i int) ( int, int) { return f2(i) }
//func f(f2 func(... interface{}) int error, i int) { f2(i) }

func main() {
//	f(fmt.Println, 3)
	r := func (a int) (int, int) { return a, a*a }
	fmt.Printf("%T\n", r)
	fmt.Println(f(r, 2))
}
