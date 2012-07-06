package main
import "fmt"
//anon fields
type tanon struct {
	int
	b int
}

func main() {
	t := tanon{1,22}
	fmt.Println(t)
}