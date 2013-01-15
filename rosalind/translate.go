package rosalind

var translationTable = map[string]string{
	"UUU": "F", "UUC": "F", "UUA": "L", "UUG": "L",
	"UCU": "S", "UCC": "S", "UCA": "S", "UCG": "S",
	"UAU": "Y", "UAC": "Y", "UAA": "", "UAG": "",
	"UGU": "C", "UGC": "C", "UGA": "", "UGG": "W",

	"CUU": "L", "CUC": "L", "CUA": "L", "CUG": "L",
	"CCU": "P", "CCC": "P", "CCA": "P", "CCG": "P",
	"CAU": "H", "CAC": "H", "CAA": "Q", "CAG": "Q",
	"CGU": "R", "CGC": "R", "CGA": "R", "CGG": "R",

	"AUU": "I", "AUC": "I", "AUA": "I", "AUG": "M",
	"ACU": "T", "ACC": "T", "ACA": "T", "ACG": "T",
	"AAU": "N", "AAC": "N", "AAA": "K", "AAG": "K",
	"AGU": "S", "AGC": "S", "AGA": "R", "AGG": "R",

	"GUU": "V", "GUC": "V", "GUA": "V", "GUG": "V",
	"GCU": "A", "GCC": "A", "GCA": "A", "GCG": "A",
	"GAU": "D", "GAC": "D", "GAA": "E", "GAG": "E",
	"GGU": "G", "GGC": "G", "GGA": "G", "GGG": "G",
}

func TranslateRNA(rna string) (protein string) {
	var triplet string
	// while the input string is 3 characters or longer
	for len(rna) >= 3 {
		triplet = rna[:3] // read first 3 chars
		// if triplet == stop codon: return
		if translationTable[triplet] == "" {
			return protein
		}
		// else look up triplet and add AA to
		// growing protein string
		protein = protein + translationTable[triplet]
		rna = rna[3:] // chop off first 3 chars of input
	}
	// run out of rna, return protein
	return protein
}
