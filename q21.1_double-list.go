package main

import ("fmt"; "container/list")

func main() {
	l := list.New()
	l.PushBack(100)
	l.PushBack(101)
	l.PushBack("hl")
	fmt.Println(l.Len())

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e, "VALUE:", e.Value)
	}
}