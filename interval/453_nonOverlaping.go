package interval

import (
	"fmt"
	"math"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) < 2 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] ||
			(intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})
	for i, in := range intervals {
		intervals[i] = append(in, intervalLen(in))
	}
	fmt.Printf("%+v\n", intervals)

	var result int
	current := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if current[1] > intervals[i][0] || (current[0] == intervals[i][0] && current[1] <= intervals[i][1]) {
			if intervalLen(current) > intervalLen(intervals[i]) {
				fmt.Printf("droping: %#v\n", current)
				current = intervals[i]
			} else {
				fmt.Printf("droping: %#v\n", intervals[i])
			}
			result++
		} else {
			current = intervals[i]
		}
	}

	return result
}

// eraseOverlapIntervals2
// we order the intervals by the end point
// we set last end value to minimum int
// for every interval on our list, we see if the interval overlaps with the last end
// if it does, we increase our erase counter
// otherwise we update our last_end to the current interval end
func eraseOverlapIntervals2(intervals [][]int) int {
	var result int
	lastEnd := math.MinInt32
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	for _, interval := range intervals {
		if interval[0] >= lastEnd {
			lastEnd = interval[1]
		} else {
			result++
		}
	}
	return result
}

func intervalLen(interval []int) int {
	return interval[1] - interval[0]
}
