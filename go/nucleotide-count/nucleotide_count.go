// Package dna - work outthe frquency count of nucleotides in a DNS strand
package dna

import (
	"errors"
)

// Histogram map - still not clear to me why rune and not string
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts - fucntion to count the nucleotides
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'G': 0, 'C': 0, 'T': 0, 'A': 0}
	for _, nucleotide := range d {
		// underscore would be the vale of h[nucleotide]. found is true if the value has a match. A good explanation is at:
		// https://yourbasic.org/golang/check-map-contains-key/
		if _, found := h[nucleotide]; found {
		} else {
			return nil, errors.New("invalid nucleotide")
		}
		h[nucleotide] = h[nucleotide] + 1
	}
	return h, nil
}
