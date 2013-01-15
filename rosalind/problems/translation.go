package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kgori/mygolang/rosalind"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no filename given")
		return
	}
	if s, err := ioutil.ReadFile(os.Args[1]); err == nil {
		fmt.Println(rosalind.Translate(string(s)))
	}
}
