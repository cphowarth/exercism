// Package triangle work out triangle type from side lengths
package triangle

import "fmt"
//import "math"

type Kind string

const (
    NaT Kind = "not a triangle"
    Equ Kind = "equilateral"
    Iso Kind = "isosceles"
    Sca Kind = "scalene"
)

// KindFromSides function to work out the triangle type from input lengths
func KindFromSides(a, b, c float64) Kind {
	//var k Kind
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
	if ( ( a==b ) && ( b==c) && ( a==c) ) {
		return true
	}
	return false
}

// IsIsosceles - determine if isosceles triange
func IsIsosceles(a, b, c float64) bool {
	if ( ( ( a==b ) && ( a!=c ) ) || ( ( b==c ) && ( a!=c ) ) || ( ( a==c ) && ( a!=b) ) ) {
		return true
	}
	return false
}

// scalene - determine if scalar triange
func IsScalene(a, b, c float64) bool {
	if ( ( a!=b ) && ( a!=c ) && (b!=c) ) {
		return true
	}
	return false
}

// IsNotATRiangle - determine if not a triange
func IsNotATriangle(a, b, c float64) bool {
	if ( ( a<=0 ) || ( b<=0 ) || (c<=0) ) {
		return true
	}
	//fmt.Println(math.Pow(a,2),math.Pow(b,2),math.Pow(c,2))
	//if ( ( (math.Pow(a,2) + math.Pow(b,2) ) != math.Pow(c,2) ) || ( (math.Pow(a,2) + math.Pow(c,2) ) != math.Pow(b,2) )  || ( math.Pow(b,2) + math.Pow(c,2) ) != math.Pow(a,2) ) {
	//if ( ( (a + b) >= c ) && ( (a + c) >= b ) && ( (b + c) >= a ) ) {
	if a + b < c || a + c < b || b + c < a  {
		return true
	}
	return false
}
