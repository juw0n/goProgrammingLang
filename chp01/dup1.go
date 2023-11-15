// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	inputs := bufio.NewScanner(os.Stdin)
	for inputs.Scan(){
		counts[inputs.Text()]++
	}
	for line, n := range counts {
		if n > 1{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
