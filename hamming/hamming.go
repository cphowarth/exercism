package hamming

import (
	"errors"
	"fmt"
)

// Distance - calcluate the number of differences beween the DNA sequences
func Distance(a, b string) (int, error) {
	fmt.Println(a)
	fmt.Println(b)
	distance := 0
	var err error
	if len(a) != len(b) {
		err = errors.New("sequence lengths do not match")
		return 0, err
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(string(a[i]), string(b[i]))
		//if a[i] != b[i]{
		if string(a[i]) != string(b[i]) {
			distance++
		}
	}
	return distance, err
}