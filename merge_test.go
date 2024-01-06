package merge

import (
	"intervals/merge/interval"
	"testing"
)

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
	merged, ok := MergeIntervalsIfOverlap(*a, *b)
	if !ok {
		t.Errorf("Expected overlap, got none")
	}

	if merged.Start != 1 {
		t.Errorf("Expected start to be 1, got %d", merged.Start)
	}

	if merged.End != 4 {
		t.Errorf("Expected end to be 4, got %d", merged.End)
	}
}

func TestMergeIntervalIfNoOverlap(t *testing.T) {
	a, _ := interval.New(1, 3)
	b, _ := interval.New(4, 6)
	merged, error := MergeIntervalsIfOverlap(*a, *b)
	if error {
		t.Errorf("Expected no overlap, got one")
	}

	if merged.Start != 0 {
		t.Errorf("Expected start to be 0, got %d", merged.Start)
	}

	if merged.End != 0 {
		t.Errorf("Expected end to be 0, got %d", merged.End)
	}
}

func TestMergeWithCollidingIntervals(t *testing.T) {

}

func TestMergeWithIntervalsWithIdenticalStartElements(t *testing.T) {
	a, _ := interval.New(3, 6)
	b, _ := interval.New(3, 12)
	c, _ := interval.New(3, 5)
	intervals := []interval.Interval{*a, *b, *c}
	merged, error := Merge(intervals)

	if error != nil {
		t.Errorf("Error while merging intervals")
	}

	println(merged)
}
