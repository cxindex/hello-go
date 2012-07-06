package main


func main() {
	a := new([9999999]int64)
	var b *[]int64 = new([]int64)
	for i := 0; i < 99999999; i++ {
		//a = new([9999999]]int64)
		*b = make([]int64, 9999999)
	}
	println(a)
	println(b)
}
