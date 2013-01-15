package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func prob(n float64) (r float64) {
	r = (n*n + (1-n)*(1-n)) / 2
	return r
}

func main() {
	var number float64
	var numbers []float64
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		return
	}
	split := strings.NewReader(strings.TrimSpace(string(data)))
	for {
		_, err := fmt.Fscanf(split, "%f", &number)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			break
		}
		numbers = append(numbers, number)
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%f ", prob(numbers[i]))
	}
	fmt.Println()
}
