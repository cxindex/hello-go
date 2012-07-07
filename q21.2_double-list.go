package main

import "fmt"

type List struct {
	I int
	L [3][3]interface{}
}

func (l *List) Push(x interface{}) {
	fmt.Println(len(l.L))
	// boooring
	// if l.I < 1 {
	// 	l.L[l.I][0] = nil
	// } else {
	// 	l.L[l.I][0] = &l.L[l.I-1][0]
	// 	fmt.Println("H", l.L[l.I-1][1])
	// }
	// if l.I >= 9 {
	// 	l.L[l.I][2] = nil
	// } else{
	// 	l.L[l.I][2] = &l.L[l.I+1][2]
	// }
	l.L[l.I][1] = x
	l.I++
}

func main() {
	l := new(List)
	l.Push(1)
	fmt.Println(l.L)
	l.Push("hello")
	fmt.Println(l.L)}
