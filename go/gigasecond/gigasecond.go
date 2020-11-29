// Package gigasecond - works our the amount of time passed after a gigasecond
package gigasecond

import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(1e9 * time.Second)
	return t
}
