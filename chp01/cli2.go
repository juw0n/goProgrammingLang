// cli2 print outbita command line argunents
package main
import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
		fmt.Println(i, ":", os.Args[i])
	}
}
