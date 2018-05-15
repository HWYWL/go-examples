package main

import (
	"strings"
	"fmt"
	"io"
)

func main() {
	r := strings.NewReader("Hello World!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("size = %v, err = %v, byte = %v\n", n, err, b)
		//b[:n] == b[0:n]
		fmt.Printf("value = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}
