package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/kgori/mygolang/rosalindFunctions"
)

func main() {

	var line1, line2 string
	var line []byte
	var prefix bool
	var err error

	if len(os.Args) < 2 {
		fmt.Println("no filename given")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)

infiniteLoop:
	for {
		line, prefix, err = reader.ReadLine()

		switch {
		case err != nil:
			if err == io.EOF {
				break infiniteLoop
			}
		case prefix:
			err = errors.New("Line too long")
			fmt.Println(string(line[:100]))
			break infiniteLoop
		case len(line) == 0:
			continue
		case line1 <= "":
			line1 = string(line)
		case line1 > "" && line2 <= "":
			line2 = string(line)
		case line1 > "" && line2 > "":
			break infiniteLoop
		default:
			err = errors.New("No lines found")
			break infiniteLoop
		}
	}
	d := rosalindFunctions.Hamming(line1, line2)
	fmt.Printf("%d\n", d)
}
