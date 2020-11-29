// Package proverb - output sequential lines of a proverb
package proverb

// Proverb - for given input list, convert to a proverb
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}
	var provout []string
	for index, value := range rhyme {
		if index == len(rhyme)-1 {
			provout = append(provout, "And all for the want of a "+rhyme[0]+".")
		} else {
			provout = append(provout, "For want of a "+value+" the "+rhyme[index+1]+" was lost.")
		}
	}
	return provout
}
