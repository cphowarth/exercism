package raindrops

import "strconv"

// Convert - function to determine if number is divisibe by 3,5 or 7
func Convert(number int) string {
	var raindrop = ""
	if number%3 == 0 {
		raindrop += "Pling"
	}
	if number%5 == 0 {
		raindrop += "Plang"
	}
	if number%7 == 0 {
		raindrop += "Plong"
	}
	if raindrop == "" {
		return strconv.Itoa(number)
	}
	return raindrop
}
