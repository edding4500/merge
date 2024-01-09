package merge

import (
	"intervals/merge/interval"
	"sync"
)

func splitIntervals(intervals []interval.Interval) ([]interval.Interval, []interval.Interval) {
	// for odd number of intervals, the left slice will be bigger
	if len(intervals)%2 == 0 {
		return intervals[:len(intervals)/2], intervals[len(intervals)/2:]
	} else {
		return intervals[:len(intervals)/2+1], intervals[len(intervals)/2+1:]
	}
}

func splitIntervalsBy(intervals []interval.Interval, by int) [][]interval.Interval {
	if by <= 0 {
		return [][]interval.Interval{intervals}
	}
	var splitIntervals [][]interval.Interval
	for i := 0; i < len(intervals); i += by {
		if i+by > len(intervals) {
			splitIntervals = append(splitIntervals, intervals[i:])
		} else {
			splitIntervals = append(splitIntervals, intervals[i:i+by])
		}
	}
	return splitIntervals
}

func MergeParallel(intervalsList [][]interval.Interval, parallelism int) ([]interval.Interval, error) {
	if parallelism <= 0 {

		combinedIntervals := []interval.Interval{}
		for _, intervals := range intervalsList {
			combinedIntervals = append(combinedIntervals, intervals...)
		}
		return Merge(combinedIntervals)
	}

	var wg sync.WaitGroup

	wg.Add(len(intervalsList))

	var mergedIntervals []interval.Interval
	var mergingErrors []error

	for _, intervals := range intervalsList {
		go func(splitInterval []interval.Interval) {
			defer wg.Done()
			merged, err := Merge(splitInterval)
			if err != nil {
				mergingErrors = append(mergingErrors, err)
			} else {
				mergedIntervals = append(mergedIntervals, merged...)
			}
		}(intervals)
	}

	wg.Wait()

	if len(mergingErrors) > 0 {
		return nil, mergingErrors[0]
	}

	return Merge(mergedIntervals)
}
