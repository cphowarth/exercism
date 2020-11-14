package main

import "fmt"

type Histogram map[string]int

func main() {
	h := Histogram{"G": 0,"C": 0}
	fmt.Println(h)
}
