package rosalind

import "fmt"

func Hamming(s1, s2 string) int {

	var d int
	ls1 := len(s1)
	ls2 := len(s2)

	if ls1 != ls2 {
		fmt.Println("string lengths do not match")
		return -1
	}

	for i, char := range s1 {
		if char != rune(s2[i]) {
			d++
		}
	}

	return d
}
