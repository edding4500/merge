package merge

import (
	"intervals/merge/interval"
	"math/rand"
	"testing"
)

func TestSplitIntervalsWithEvenNumberOfIntervals(t *testing.T) {
	i1, _ := interval.New(1, 5)
	i2, _ := interval.New(13, 18)
	i3, _ := interval.New(4, 8)
	i4, _ := interval.New(3, 15)

	intervals := []interval.Interval{*i1, *i2, *i3, *i4}

	left, right := splitIntervals(intervals)

	if len(left) != 2 {
		t.Errorf("Expected 2 intervals, got %d", len(left))
	}

	if len(right) != 2 {
		t.Errorf("Expected 2 intervals, got %d", len(right))
	}
}

func TestSplitIntervalsByWithEvenNumberOfIntervals(t *testing.T) {
	i1, _ := interval.New(1, 5)
	i2, _ := interval.New(13, 18)
	i3, _ := interval.New(4, 8)
	i4, _ := interval.New(3, 15)

	intervals := []interval.Interval{*i1, *i2, *i3, *i4}

	splitIntervals := splitIntervalsBy(intervals, 2)

	if len(splitIntervals) != 2 {
		t.Errorf("Expected 2 intervals, got %d", len(splitIntervals))
	}

	if len(splitIntervals[0]) != 2 {
		t.Errorf("Expected 2 intervals, got %d", len(splitIntervals[0]))
	}

	if len(splitIntervals[1]) != 2 {
		t.Errorf("Expected 2 intervals, got %d", len(splitIntervals[1]))
	}
}

func TestSplitIntervalsWithOddNumberOfIntervals(t *testing.T) {
	i1, _ := interval.New(1, 5)
	i2, _ := interval.New(13, 18)
	i3, _ := interval.New(4, 8)
	i4, _ := interval.New(3, 15)
	i5, _ := interval.New(23, 80)

	intervals := []interval.Interval{*i1, *i2, *i3, *i4, *i5}

	left, right := splitIntervals(intervals)

	if len(left) != 3 {
		t.Errorf("Expected 3 intervals, got %d", len(left))
	}

	if len(right) != 2 {
		t.Errorf("Expected 2 intervals, got %d", len(right))
	}
}

func TestSplitIntervalsByWithOddNumberOfIntervals(t *testing.T) {
	i1, _ := interval.New(1, 5)
	i2, _ := interval.New(13, 18)
	i3, _ := interval.New(4, 8)
	i4, _ := interval.New(3, 15)
	i5, _ := interval.New(23, 80)

	intervals := []interval.Interval{*i1, *i2, *i3, *i4, *i5}

	splitIntervals := splitIntervalsBy(intervals, 2)

	if len(splitIntervals) != 3 {
		t.Errorf("Expected 3 intervals, got %d", len(splitIntervals))
	}

	if len(splitIntervals[0]) != 2 {
		t.Errorf("Expected 2 intervals, got %d", len(splitIntervals[0]))
	}

	if len(splitIntervals[1]) != 2 {
		t.Errorf("Expected 2 intervals, got %d", len(splitIntervals[1]))
	}

	if len(splitIntervals[2]) != 1 {
		t.Errorf("Expected 1 intervals, got %d", len(splitIntervals[2]))
	}
}

func TestSplitIntervalsWithEmptyIntervals(t *testing.T) {
	intervals := []interval.Interval{}

	left, right := splitIntervals(intervals)

	if len(left) != 0 {
		t.Errorf("Expected 0 intervals, got %d", len(left))
	}

	if len(right) != 0 {
		t.Errorf("Expected 0 intervals, got %d", len(right))
	}
}

func TestSplitIntervalsByWithEmptyIntervals(t *testing.T) {
	intervals := []interval.Interval{}

	splitIntervals := splitIntervalsBy(intervals, 2)

	if len(splitIntervals) != 0 {
		t.Errorf("Expected 0 intervals, got %d", len(splitIntervals))
	}
}

func TestSplitIntervalsByByOneWithEmptyIntervals(t *testing.T) {
	intervals := []interval.Interval{}

	splitIntervals := splitIntervalsBy(intervals, 1)

	if len(splitIntervals) != 0 {
		t.Errorf("Expected 0 intervals, got %d", len(splitIntervals))
	}
}

func TestSplitIntervalsByByOneWithOneInterval(t *testing.T) {
	i1, _ := interval.New(1, 5)

	intervals := []interval.Interval{*i1}

	splitIntervals := splitIntervalsBy(intervals, 1)

	if len(splitIntervals) != 1 {
		t.Errorf("Expected 1 intervals, got %d", len(splitIntervals))
	}

	if len(splitIntervals[0]) != 1 {
		t.Errorf("Expected 1 intervals, got %d", len(splitIntervals[0]))
	}
}

func TestSplitIntervalsByByTwoWithOneInterval(t *testing.T) {
	i1, _ := interval.New(1, 5)

	intervals := []interval.Interval{*i1}

	splitIntervals := splitIntervalsBy(intervals, 2)

	if len(splitIntervals) != 1 {
		t.Errorf("Expected 1 intervals, got %d", len(splitIntervals))
	}

	if len(splitIntervals[0]) != 1 {
		t.Errorf("Expected 1 intervals, got %d", len(splitIntervals[0]))
	}
}

func TestSplitInvervalsByByThreeWithTwoIntervals(t *testing.T) {
	i1, _ := interval.New(1, 5)
	i2, _ := interval.New(13, 18)

	intervals := []interval.Interval{*i1, *i2}

	splitIntervals := splitIntervalsBy(intervals, 3)

	if len(splitIntervals) != 1 {
		t.Errorf("Expected 1 intervals, got %d", len(splitIntervals))
	}

	if len(splitIntervals[0]) != 2 {
		t.Errorf("Expected 2 intervals, got %d", len(splitIntervals[0]))
	}
}

func TestSplitIntervalsByByThreeWithSixIntervals(t *testing.T) {
	i1, _ := interval.New(1, 5)
	i2, _ := interval.New(13, 18)
	i3, _ := interval.New(4, 8)
	i4, _ := interval.New(3, 15)
	i5, _ := interval.New(23, 80)
	i6, _ := interval.New(23, 80)

	intervals := []interval.Interval{*i1, *i2, *i3, *i4, *i5, *i6}

	splitIntervals := splitIntervalsBy(intervals, 3)

	if len(splitIntervals) != 2 {
		t.Errorf("Expected 2 intervals, got %d", len(splitIntervals))
	}

	if len(splitIntervals[0]) != 3 {
		t.Errorf("Expected 3 intervals, got %d", len(splitIntervals[0]))
	}

	if len(splitIntervals[1]) != 3 {
		t.Errorf("Expected 3 intervals, got %d", len(splitIntervals[1]))
	}
}

func TestSplitIntervalsByByZero(t *testing.T) {
	i1, _ := interval.New(1, 5)
	i2, _ := interval.New(13, 18)
	i3, _ := interval.New(4, 8)
	i4, _ := interval.New(3, 15)
	i5, _ := interval.New(23, 80)
	i6, _ := interval.New(23, 80)

	intervals := []interval.Interval{*i1, *i2, *i3, *i4, *i5, *i6}

	splitIntervals := splitIntervalsBy(intervals, 0)

	if len(splitIntervals) != 1 {
		t.Errorf("Expected 1 interval, got %d", len(splitIntervals))
	}

	if len(splitIntervals[0]) != 6 {
		t.Errorf("Expected 6 intervals, got %d", len(splitIntervals[0]))
	}

	if splitIntervals[0][0].Start != 1 {
		t.Errorf("Expected start to be 1, got %d", splitIntervals[0][0].Start)
	}

	if splitIntervals[0][0].End != 5 {
		t.Errorf("Expected end to be 5, got %d", splitIntervals[0][0].End)
	}
}

func TestMergeParallelWithEvenNumberOfIntervals(t *testing.T) {
	i1, _ := interval.New(1, 5)
	i2, _ := interval.New(13, 18)
	i3, _ := interval.New(4, 8)
	i4, _ := interval.New(3, 15)

	intervals := []interval.Interval{*i1, *i2, *i3, *i4}

	splittedIntervals := splitIntervalsBy(intervals, 2)

	merged, err := MergeParallel(splittedIntervals, 2)

	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	if len(merged) != 1 {
		t.Errorf("Expected 1 interval, got %d", len(merged))
	}

	if merged[0].Start != 1 {
		t.Errorf("Expected start to be 1, got %d", merged[0].Start)
	}

	if merged[0].End != 18 {
		t.Errorf("Expected end to be 18, got %d", merged[0].End)
	}
}

func BenchmarkMergeParallelWithEvenNumberOfIntervals(b *testing.B) {

	var intervals []interval.Interval

	for i := 0; i < 1000000; i++ {
		randStart := rand.Intn(1000000)
		randEnd := randStart + rand.Intn(100)
		intervals = append(intervals, interval.Interval{Start: randStart, End: randEnd})
	}

	parallelism := 16

	splittedIntervals := splitIntervalsBy(intervals, parallelism)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		MergeParallel(splittedIntervals, parallelism)
	}
}

func BenchmarkMergeWithEvenNumberOfIntervals(b *testing.B) {

	var intervals []interval.Interval

	for i := 0; i < 1000000; i++ {
		randStart := rand.Intn(1000000)
		randEnd := randStart + rand.Intn(100)
		intervals = append(intervals, interval.Interval{Start: randStart, End: randEnd})
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		Merge(intervals)
	}
}

func TestMergingInGoRoutines(t *testing.T) {
	i1, _ := interval.New(1, 5)
	i2, _ := interval.New(13, 18)
	i3, _ := interval.New(4, 8)
	i4, _ := interval.New(3, 15)
	i5, _ := interval.New(23, 80)

	intervals := []interval.Interval{*i1, *i2, *i3, *i4, *i5}

	left, right := splitIntervals(intervals)

	go Merge(left)
	go Merge(right)

	leftMerged, err1 := Merge(left)
	if err1 != nil {
		t.Errorf("Expected no error, got %s", err1)
	}

	rightMerged, err2 := Merge(right)
	if err2 != nil {
		t.Errorf("Expected no error, got %s", err2)
	}

	merged, err3 := Merge(append(leftMerged, rightMerged...))
	if err3 != nil {
		t.Errorf("Expected no error, got %s", err3)
	}

	if len(merged) != 2 {
		t.Errorf("Expected 2 intervals, got %d", len(merged))
	}

}
