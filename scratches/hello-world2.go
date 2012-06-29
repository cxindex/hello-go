package main // #26 and #27 and #28
//pointers, struct literals, new func
import "fmt"

type vertex struct {X, Y int}

var (
	p_ = vertex{1, 2}  // has type Vertex
	q_ = &vertex{1, 2} // has type *Vertex
	r_ = vertex{X: 1}  // Y:0 is implicit
	s_ = vertex{}      // X:0 and Y:0
)


func main() {
	p := vertex{1,2}
	q := &p
	q.X = 1e9
	fmt.Println(p)
	fmt.Printf("%T %T\n", q, p)

	//for maximum obviously
	var z int = 1
	var zp *int = &z;
	*zp = *zp + 101
	fmt.Printf("%T %T %v %v. addr: %v\n", z, zp, z, *zp, zp)

	fmt.Println(p_, q_, r_, s_)

	v := new (vertex) //or var v *vertex = new (vertext)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
}


