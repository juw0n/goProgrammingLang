// Fetch prints the content found at a URL.
// Use it instead of ioutil.ReadAll to copy the response body to os.Stdout wit// hout requiring a buffer large enough to hold the entire stream. 


package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)
func main() {
	for _, url := range os.Args[1:] {
		resp, err:= http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()

		b, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

