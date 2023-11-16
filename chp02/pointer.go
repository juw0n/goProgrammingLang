// comparing pointers

package main
import "fmt"

var p = f()
func f() *int {
	v := 1
	return &v
}

func main() {
	fmt.Println(f() == f())
	fmt.Println(f())
}
