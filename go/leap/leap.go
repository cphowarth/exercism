// Package leap - work out if it is a leap year
package leap

// IsLeapYear function to determine if leap year
func IsLeapYear(year int) bool {
	isleapyear := false
	if year%4 == 0 {
		if year%100 != 0 {
			isleapyear = true
		} else {
			isleapyear = false
			if year%400 == 0 {
				isleapyear = true
			}
		}
	}
	return isleapyear
}
