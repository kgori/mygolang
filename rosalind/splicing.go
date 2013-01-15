package rosalind

import "bytes"

func Splice(gene, intron []byte) []byte {
	return bytes.Replace(gene, intron, []byte{}, 1)
}
