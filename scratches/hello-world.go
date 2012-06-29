package main
import (
	"fmt"
	"math"
	"math/cmplx"
)

const Pi = 3.14
const (
	Big = 1<<100
	Small = Big>>99
)

type Vertex struct {
	X, Y int
}
	

//	The var statement declares a list of variables; as in function argument lists, the type is last.
var x, y, z int = 1, 2, 3
var c, python, java = true, false, "нет!"
var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 -1
	zz complex128 = cmplx.Sqrt(-5+12i)
)

func needInt(x int) int { return x*10+1 }
func needFloat(x float64) float64 { return x*0.1 }
// When two or more consecutive named function parameters share a type, you can omit the type from all but the last. 
// func add(x int, y int) int {
func add(x, y int) int { return x + y }
func swap(x, y string) (string, string) { return y, x }
func split(sum int) (x, y int) {
	x = sum * 4/9
	y = sum - x
	return
	// wil return multi. i.e. by default will be returned x and y
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i" //imaginary number
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 { // or (v float64) for default and all_func_vision
	//Variables declared by the statement are only in scope until the end of the if.
	if v := math.Pow(x, n); v < lim {
	 	return v
	} else {
		fmt.Printf("From pow: %g >= %g\n", v, lim)
	}
	//cant use v here, though
	return lim
}


func main() {
	fmt.Printf("Hello, World! And this is type of number 4: %T\n Тестовое Пи: %v\n And now about problems: %g\n\n",
		4,
		math.Pi,
		math.Nextafter(2,3))

	fmt.Println(Small)
	fmt.Println("Test for Println and add:", add(2,8))
	a, b := swap("hello", "world!") //or  b, a := "hello", "world"
	fmt.Println("Test for Swap:", a, b)	
	fmt.Printf("Test for Split: "); fmt.Println(split(17))

	fmt.Println("Global var:", x, y, z, c, python, java, "\nGlobal const Pi:", Pi)
	//Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type
	//Outside a function, every construct begins with a keyword and the := construct is not available
	x, y, z := 10, 11, 12
	fmt.Println("Local var:", x, y, z)

	fmt.Printf("needInt(Small): %v\tneedFloat(Small): %v\n", needInt(Small), needFloat(Small))
	fmt.Printf("needInt(Small): %v\tneedFloat(Big): %v\n", needInt(Small), needFloat(Big)) //nI Big - overflows

	var sum float64
	addv := 100
	for i := 0.0; i <= 5; i++ { // or without pre and post: for ; i <= 5; or : for i <= 5; for ; ; - forever; for - same
		sum += i; addv += 100
		fmt.Printf("I: %v, addv: %v, from add(): %v, sum: %v\n", i, addv, add(10, addv), sum)
	}

	fmt.Printf("Sqrt(2): %v\t Sqrt(-4): %v\n", sqrt(2), sqrt(-4))
	fmt.Println(
		"Pow(3, 2, 10):", pow(3, 2, 10),
		"\tPow(3, 3, 20):", pow(3, 3, 10),
	)

	const f = "Type: %T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, zz, zz)

	v := Vertex{1,2}
	v.X = 4
	fmt.Printf(f, v.X, v.X)
	fmt.Printf(f, v.Y, v.Y)
	
}
