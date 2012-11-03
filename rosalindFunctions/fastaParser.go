package rosalindFunctions

// package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type sequence struct {
	ID  string
	Seq string
}

// FASTAReader type for reading FASTA files.
type FastaReader struct {
	reader *bufio.Reader
	nextID string
}

func NewFastaReader(newBufioReader *bufio.Reader) FastaReader {
	return FastaReader{reader: newBufioReader}
}

func NewSequence() sequence {
	return sequence{
		ID:  "",
		Seq: "",
	}
}

func (fastareader *FastaReader) ParseFastaFile() (item sequence, err error) {
	var line, seq []byte
	var prefix bool
	id := fastareader.nextID
	fastareader.nextID = ""
reading:
	for {
		line, prefix, err = fastareader.reader.ReadLine()
		switch {
		case err != nil:
			if err == io.EOF && id > "" {
				err = nil
			}
			break reading
		case prefix:
			err = errors.New("Line too long")
			fmt.Println(string(line[:100]))
			break reading
		case len(line) == 0:
			continue // skip empty line
		case line[0] == '>':
			if id > "" {
				fastareader.nextID = string(line[1:])
				break reading
			}
			id = string(line[1:])
		case id > "":
			seq = append(seq, line...)
		default:
			// ignore
		}
	}
	item.ID = id
	item.Seq = string(seq)
	return item, err
}

func GCContent(seq string) float64 {
	gContent := float64(strings.Count(seq, "G"))
	cContent := float64(strings.Count(seq, "C"))
	seqLength := float64(len(seq))
	gcContent := (cContent + gContent) / seqLength
	return gcContent
}
