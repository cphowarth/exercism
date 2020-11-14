// Package triangle work out triangle type from side lengths
package triangle

import "fmt"
import "math"

// Kind string - not immediately clear to me why the type is required
type Kind string

// Define strings to be returned from the function KindFromSides
const (
	NaT Kind = "not a triangle"
	Equ Kind = "equilateral"
	Iso Kind = "isosceles"
	Sca Kind = "scalene"
)

// KindFromSides function to work out the triangle type from input lengths
func KindFromSides(a, b, c float64) Kind {
	fmt.Println(a, b, c)
	if IsNotATriangle(a, b, c) == true {
		// fmt.Println("IsNotATriangle")
		return NaT
	}
	if IsEquilateral(a, b, c) == true {
		// fmt.Println("IsEquilateral")
		return Equ
	}
	if IsIsosceles(a, b, c) == true {
		// fmt.Println("IsIsosceles")
		return Iso
	}
	if IsScalene(a, b, c) == true {
		// fmt.Println("IsScalene")
		return Sca
	}
	return "should never get here"
}

// IsEquilateral - determine if equilateral triange
func IsEquilateral(a, b, c float64) bool {
	if (a == b) && (b == c) && (a == c) {
		return true
	}
	return false
}

// IsIsosceles - determine if isosceles triange
func IsIsosceles(a, b, c float64) bool {
	if ((a == b) && (a != c)) || ((b == c) && (a != c)) || ((a == c) && (a != b)) {
		return true
	}
	return false
}

// IsScalene - determine if scalar triange
func IsScalene(a, b, c float64) bool {
	if (a != b) && (a != c) && (b != c) {
		return true
	}
	return false
}

// IsNotATriangle - determine if not a triange
func IsNotATriangle(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return true
	}
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return true
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return true
	}
	if a+b < c || a+c < b || b+c < a {
		return true
	}
	return false
}
