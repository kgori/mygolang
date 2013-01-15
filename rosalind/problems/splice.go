package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/kgori/mygolang/rosalind"
	"io"
	"os"
	// "sort"
)

type ByLength [][]byte

// Implement sorting by length for [][]byte
func (b ByLength) Len() int {
	return len(b)
}

func (b ByLength) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByLength) Less(i, j int) bool {
	return len(b[i]) > len(b[j])
}

func main() {
	var lines [][]byte
	var line []byte
	var prefix bool

	f, err := os.Open(os.Args[1]) // open file
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

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
			lines = append(lines, line)
		}
	}

	gene := lines[0]
	introns := lines[1:]

	// sort.Sort(ByLength(introns))

	for _, intron := range introns {
		gene = rosalind.Splice(gene, intron)
	}

	asRNA := string(bytes.Replace(gene, []byte{'T'}, []byte{'U'}, -1))
	fmt.Printf("%s\n", rosalind.Translate(asRNA))
}
