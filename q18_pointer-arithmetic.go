package main

import "fmt"

type st struct { int }

func main() {
	i := 4
	p := &i
	fmt.Println(i, *p)
	*p++; (*p)++
	fmt.Println(i, *p)


	f := float64(i) + 0.1; pf := &f
	fmt.Println(f, *pf)
	*pf++; (*pf)++
	fmt.Println(f, *pf)
	
	st := st{1}; stf := &st
	fmt.Println(st, *stf)
//	*stf++;	(*stf)++ //wont work. i,e, only numerical values

	s := "hello"; sp := &s
//	*s++ //same here
	fmt.Println(s, *sp)
}