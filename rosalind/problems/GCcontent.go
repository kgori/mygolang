package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/kgori/mygolang/rosalind"
)

func main() {
	var currGC float64
	var bestID string
	bestGC := 0.0
	item := rosalind.NewSequence()

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	myFastaReader := rosalind.NewFastaReader(bufio.NewReader(f))

infiniteLoop:
	for {
		item, err = myFastaReader.ParseFastaFile()
		switch {
		case err == io.EOF:

			break infiniteLoop

		default:
			currGC = rosalind.GCContent(item.Seq)
			if currGC > bestGC {
				bestGC = currGC
				bestID = item.ID
			}
		}
	}
	fmt.Println(bestID)
	fmt.Printf("%.3f\n", bestGC*100)
}
