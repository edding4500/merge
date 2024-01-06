package interval

import "errors"

type Interval struct {
	Start int
	End   int
}

func New(start, end int) (*Interval, error) {
	if start > end {
		// throw error
		return nil, errors.New("Start cannot be greater than end")
	} else {
		return &Interval{start, end}, nil
	}
}

func (i Interval) Overlaps(j Interval) bool {
	return i.Start <= j.End && j.Start <= i.End
}
