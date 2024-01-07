package interval

import "errors"

//
// Represents an interval with a start and end.
type Interval struct {
	Start int
	End   int
}

//
// Constructs a new interval with start and end.
// If start > end, an error is returned.
func New(start, end int) (*Interval, error) {
	if start > end {
		// throw error
		return nil, errors.New("Start cannot be greater than end")
	} else {
		return &Interval{start, end}, nil
	}
}

// Returns true if the interval overlaps with the given interval.
//
// Example:
// 	a, _ := interval.New(1, 3)
// 	b, _ := interval.New(2, 4)
// 	a.Overlaps(b) // true
// 	a, _ := interval.New(1, 3)
// 	b, _ := interval.New(4, 6)
// 	a.Overlaps(b) // false
func (i Interval) Overlaps(j Interval) bool {
	return i.Start <= j.End && j.Start <= i.End
}
