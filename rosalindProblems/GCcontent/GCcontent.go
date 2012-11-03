package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/kgori/mygolang/rosalindFunctions"
)

func main() {
	var currGC float64
	var bestID string
	bestGC := 0.0
	item := rosalindFunctions.NewSequence()

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	myFastaReader := rosalindFunctions.NewFastaReader(bufio.NewReader(f))

infiniteLoop:
	for {
		item, err = myFastaReader.ParseFastaFile()
		switch {
		case err == io.EOF:

			break infiniteLoop

		default:
			currGC = rosalindFunctions.GCContent(item.Seq)
			if currGC > bestGC {
				bestGC = currGC
				bestID = item.ID
			}
		}
	}
	fmt.Println(bestID)
	fmt.Printf("%.3f\n", bestGC*100)
}
