// Print the name of command
// Example: go run .\ex1.1_echo.go 1 2 3
// Result: [...]/ex1.1_echo.exe 1 2 3

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
