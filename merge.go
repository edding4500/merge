package merge

import (
	"errors"
	"intervals/merge/interval"
	"sort"
)

/**
*
*	`Merge` takes a list of intervals and returns a list of
*	merged intervals, i.e. a list of intervals where all overlapping
*	intervals are merged together. In case of an error,
* the returned error object reflects the error and is not nil.
*
* @intervals: The list of intervals to be merged
* @return: A list of merged intervals and an error object.
*
 */
func Merge(intervals []interval.Interval) ([]interval.Interval, error) {
	if len(intervals) == 0 {
		return intervals, nil
	}

	/**
	*	Intervals are sorted ascendingly by their Start element.
	* We do so to make the next step of merging work: we only
	* have to iterate once over all intervals and no longer
	* have to compare each interval with each other interval.
	* This sorting step thus avoids quadratic runtime.
	*	This step introduces **invariant A**: no remaining intervals
	* after the current interval can be merged with the last interval.
	*
	* Notes for discussion: sort.Slice is mentioned to be not stable.
	* Do we need stable sorting here? No, we dont, because the
	* order of the second element does not matter for merging.
	 */
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a].Start < intervals[b].Start
	})

	/**
	*	We keep list of all merged intervals. This adds to the
	* memory footprint of the algorithm: in the worst case for
	* an input if size n we need twice the memory. We could
	* improve here by merging in place or dropping the merged
	* intervals from the list of intervals. For now, I keep
	* it simple.
	 */
	merged := []interval.Interval{intervals[0]}
	for _, currentInterval := range intervals[1:] {
		last := merged[len(merged)-1]
		/**
		*	If the current interval overlaps with the last,
		*	we can safely merge them and update the last interval
		* to the merged one. This step is crucial to avoid quadratic
		* runtime: in the next iteration of the for loop, we compare
		* with the now updated, merged interval.
		 */
		if last.Overlaps(currentInterval) {
			mergedInverval, err := MergeInterval(last, currentInterval)
			if err != nil {
				return nil, err
			}
			merged[len(merged)-1] = *mergedInverval
		} else {
			/**
			*	If the current interval does not overlap with the last one,
			* we can move on to comparing the current interval to the remaining
			* ones. We do so by updating the merged intervals list, appending
			* the current interval. In the next iteration, we pick this very
			* element as the one for comparison to the remaing intervals.
			* We can safely do so due to invariant A: no remaining intervals
			* after the current interval can be merged with the last interval.
			* Their Start elements are bigger than the End element of the current
			* interval. If not, the current interval would have been merged with
			* with the last one.
			 */
			merged = append(merged, currentInterval)
		}
	}

	return merged, nil
}

/**
*
*	`MergeInterval` takes two intervals and returns a merged interval
*	if the two intervals overlap. In case of an error, the returned
*	error object reflects the error and is not nil.
*
* @a: The first interval
* @b: The second interval
* @return: A merged interval and an error object.
*
 */
func MergeInterval(a, b interval.Interval) (*interval.Interval, error) {
	return interval.New(
		min(a.Start, b.Start),
		max(a.End, b.End),
	)
}

// Takes two intervals and returns a merged interval
// if the two intervals overlap.
//
// In case of an error, the returned
// error object reflects the error and is not nil.
//
// If the intervals do not overlap, the returned
// error object reflects the error and is not nil.
// The returned interval is empty.
func MergeIntervalsIfOverlap(a, b interval.Interval) (interval.Interval, error) {
	if a.Overlaps(b) {
		merged, err := MergeInterval(a, b)
		return *merged, err
	}

	return interval.Interval{}, errors.New("Intervals do not overlap")
}
