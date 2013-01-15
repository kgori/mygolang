package main

import (
	"fmt"
	"github.com/kgori/mygolang/rosalind"
	"io/ioutil"
	"os"
)

func main() {
	// Read file from commandline argument [1]
	// Print reverse complement - see reverseComplement.go
	if s, err := ioutil.ReadFile(os.Args[1]); err == nil {
		fmt.Println(rosalind.ReverseComplement(string(s)))
	}
}
