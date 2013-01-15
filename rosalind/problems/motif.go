package main

import (
	"bufio"
	"fmt"
	"github.com/kgori/mygolang/rosalind"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1]) // open file
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f) // read file
	var s, t []byte
	var prefix bool
	s, prefix, err = r.ReadLine()
	t, prefix, err = r.ReadLine()

	if prefix {
		panic("Line too long")
	}
	answer := rosalind.MatchMotif(s, t)
	l := len(answer)
	for i, a := range answer {
		if i <= (l - 2) {
			fmt.Printf("%d ", a)
		} else {
			fmt.Printf("%d\n", a)
		}
	}
}
