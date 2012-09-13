//Example of a main package from 
//http://golang.org/doc/code.html#tmp_13
package main

import (
	"fmt"
	"github.com/kgori/mygolang/newmath"
)

func main() {
	fmt.Printf("Hello, World. Sqrt(2) = %v\n", newmath.Sqrt(2))
	fmt.Printf("(Example from http://golang.org/doc/code.html#tmp_13)\n")
}
