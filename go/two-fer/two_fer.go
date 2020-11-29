package twofer

import "fmt"
// ShareWith - Respond with a user in the sentence two for you....
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	return fmt.Sprintf("One for %s, one for me.", name)
}
