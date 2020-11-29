// Package strand - convert DNA to RNA
package strand

import "strings"

// ToRNA - convert to RNA
func ToRNA(dna string) string {
	rna := strings.Split(dna, "")
	for index, nucleotide := range rna {
		switch nucleotide {
		case "G":
			rna[index] = "C"
		case "C":
			rna[index] = "G"
		case "T":
			rna[index] = "A"
		case "A":
			rna[index] = "U"
		}
	}
	return strings.Join(rna, "")
}
