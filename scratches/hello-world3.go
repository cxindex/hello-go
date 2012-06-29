//maps. #29
package main

import "fmt"

type vrtx struct {
	lat, long float64
}

var m map[string]vrtx

func main() {
	m = make(map[string]vrtx)
	m["Bell Labs"] = vrtx{
		40.1111, 74.1111,
	}
	fmt.Println(m["Bell Labs"])


	//index with int
	m2 := make(map[int]vrtx) //new map-variable with initialization
	m2[42] = vrtx{
		140.1111, 174.1111,
	}
	fmt.Println(m2[42])
	// няшнота!
	m2[43] = vrtx{
		143.1111, 177.1111,
	}
	fmt.Println(m2[43])
	
}