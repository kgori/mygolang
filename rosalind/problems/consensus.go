package main

import (
	"bufio"
	"fmt"
	"github.com/kgori/mygolang/rosalind"
	"io"
	"os"
)

func main() {

	var prefix bool
	var line []byte

	f, err := os.Open(os.Args[1]) // open file
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	nc := rosalind.NewConsM(1)

reading:
	for {
		line, prefix, err = reader.ReadLine()
		switch {
		case err == io.EOF:
			break reading
		case prefix:
			fmt.Println(line[:100])
			panic("Line too long")
		case len(line) == 0:
			continue // skip empty line
		default:
			nc.AddBytes(line)
		}
	}

	fmt.Println(string(nc.Consensus()))
	fmt.Println(nc)

}
