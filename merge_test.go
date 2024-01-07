package merge

import (
	"intervals/merge/interval"
	"testing"
)

// Test if the intervals are merged using
// hard coded intervals.
func TestMergeWithFixedIntervals(t *testing.T) {
	i1, _ := interval.New(1, 5)
	i2, _ := interval.New(13, 18)
	i3, _ := interval.New(4, 8)
	i4, _ := interval.New(3, 15)
	i5, _ := interval.New(23, 80)

	intervals := []interval.Interval{*i1, *i2, *i3, *i4, *i5}

	merged, err := Merge(intervals)

	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	if len(merged) != 2 {
		t.Errorf("Expected 2 intervals, got %d", len(merged))
	}

	if merged[0].Start != 1 {
		t.Errorf("Expected start to be 1, got %d", merged[0].Start)
	}

	if merged[0].End != 18 {
		t.Errorf("Expected end to be 18, got %d", merged[0].End)
	}

	if merged[1].Start != 23 {
		t.Errorf("Expected start to be 23, got %d", merged[1].Start)
	}

	if merged[1].End != 80 {
		t.Errorf("Expected end to be 80, got %d", merged[1].End)
	}
}

// Test if the intervals are merged when theire
// Start elements are sorted in descending order.
func TestMergeIntervalWithInvertedIntervals(t *testing.T) {
	a, _ := interval.New(1, 3)
	b, _ := interval.New(4, 6)
	c, _ := interval.New(7, 9)
	intervals := []interval.Interval{*c, *b, *a}
	merged, err := Merge(intervals)

	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	if len(merged) != 3 {
		t.Errorf("Expected 3 intervals, got %d", len(merged))
	}
}

func TestMergeIntervalIfOverlap(t *testing.T) {
	a, _ := interval.New(1, 3)
	b, _ := interval.New(2, 4)
	merged, error := MergeIntervalsIfOverlap(*a, *b)
	if error != nil {
		t.Errorf("Expected overlap but merging failed")
	}

	if merged.Start != 1 {
		t.Errorf("Expected start to be 1, got %d", merged.Start)
	}

	if merged.End != 4 {
		t.Errorf("Expected end to be 4, got %d", merged.End)
	}
}

// Test if the intervals are not merged when they are
// not overlapping.
func TestMergeIntervalIfNoOverlap(t *testing.T) {
	a, _ := interval.New(1, 3)
	b, _ := interval.New(4, 6)
	_, error := MergeIntervalsIfOverlap(*a, *b)
	if error == nil {
		t.Errorf("Expected no overlap but merging succeeded")
	}
}

// Test if intervals are merged when they are identical
func TestMergeWithCollidingIntervals(t *testing.T) {
	a, _ := interval.New(1, 3)
	b, _ := interval.New(1, 3)
	c, _ := interval.New(1, 3)
	intervals := []interval.Interval{*a, *b, *c}
	merged, error := Merge(intervals)

	if error != nil {
		t.Errorf("Error while merging intervals")
	}

	if len(merged) != 1 {
		t.Errorf("Expected 1 interval, got %d", len(merged))
	}

	if merged[0].Start != 1 {
		t.Errorf("Expected start to be 1, got %d", merged[0].Start)
	}

	if merged[0].End != 3 {
		t.Errorf("Expected end to be 3, got %d", merged[0].End)
	}
}

// Test if the intervals are merged when they have
// identical start elements
func TestMergeWithIntervalsWithIdenticalStartElements(t *testing.T) {
	a, _ := interval.New(3, 6)
	b, _ := interval.New(3, 12)
	c, _ := interval.New(3, 5)
	intervals := []interval.Interval{*a, *b, *c}
	merged, error := Merge(intervals)

	if error != nil {
		t.Errorf("Error while merging intervals")
	}

	if len(merged) != 1 {
		t.Errorf("Expected 1 interval, got %d", len(merged))
	}
}

// Test merging of intervals with negative numbers
func TestMergeWithIntervalsWithNegativeNumbersAndIdenticalStartValue(t *testing.T) {
	a, _ := interval.New(-3, 6)
	b, _ := interval.New(-3, 12)
	c, _ := interval.New(-3, 5)
	intervals := []interval.Interval{*a, *b, *c}
	merged, error := Merge(intervals)

	if error != nil {
		t.Errorf("Error while merging intervals")
	}

	if len(merged) != 1 {
		t.Errorf("Expected 1 interval, got %d", len(merged))
	}
}

func TestMergeWithIntervalsWithNegativeNumbers(t *testing.T) {
	a, _ := interval.New(2, 6)
	b, _ := interval.New(-3, 2)
	c, _ := interval.New(-12, 5)
	intervals := []interval.Interval{*a, *b, *c}
	merged, error := Merge(intervals)

	if error != nil {
		t.Errorf("Error while merging intervals")
	}

	if len(merged) != 1 {
		t.Errorf("Expected 1 interval, got %d", len(merged))
	}
}

// test merging of a single interval
func TestMergeWithSingleInterval(t *testing.T) {
	a, _ := interval.New(1, 3)
	intervals := []interval.Interval{*a}
	merged, error := Merge(intervals)

	if error != nil {
		t.Errorf("Error while merging intervals")
	}

	if len(merged) != 1 {
		t.Errorf("Expected 1 interval, got %d", len(merged))
	}
}

// Test merging of an empty list of intervals
func TestMergeWithEmptyIntervals(t *testing.T) {
	intervals := []interval.Interval{}
	merged, error := Merge(intervals)

	if error != nil {
		t.Errorf("Error while merging intervals")
	}

	if len(merged) != 0 {
		t.Errorf("Expected 0 intervals, got %d", len(merged))
	}
}
