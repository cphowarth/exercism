// Package dna - work outthe frquency count of nucleotides in a DNS strand
package dna

import (
	"strings"
	"errors"
)

type Histogram map[string]int {
	"G": 0,
	"C": 0,
	"T": 0,
	"A": 0
}
// DNA is a list of nucleotides. Choose a suitable data type.
type DNA byte

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.

func (d DNA) Counts() (Histogram, error) {
	var h Histogram 
	var err error
        rna := strings.Split(dna, "")
	for index, nucleotide := range rna {
		if ! h[nucleotide] {
			return err = errors.New("invalid nucleotide")
		}
		h[nucleotide] = h[nucleotide] +1
	}
	return h, err
}
