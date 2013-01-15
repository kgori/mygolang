package rosalind

func ReverseComplement(s string) string {

	var complement = map[rune]rune{
		'A': 'T',
		'C': 'G',
		'G': 'C',
		'T': 'A',
	}

	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = complement[rune]
	}
	return string(runes[n:])
}
