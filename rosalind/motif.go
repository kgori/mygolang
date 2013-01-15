package rosalind

import (
	"bytes"
)

func MatchMotif(s, t []byte) []int {
	r := []int{}
	var i int
	for i < len(s) {
		if bytes.HasPrefix(s[i:], t) {
			r = append(r, i+1)
		}
		i++
	}
	return r
}
