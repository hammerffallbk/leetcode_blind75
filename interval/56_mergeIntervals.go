package interval

import "sort"

// merge
// we order the intervals on ascending order on the initial point
// we take the first interval(current) and look if the next on the list overlaps
// take the min(initial point) and max(endpoint) for the current interval
// if the current interval does not overlap with the index pointed, we added to the result list
// last current value should be added to the list as well
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0, len(intervals))
	current := []int{intervals[0][0], intervals[0][1]}
	for i := 1; i < len(intervals); i++ {
		ithInterval := []int{intervals[i][0], intervals[i][1]}
		if current[1] >= ithInterval[0] {
			current[0] = min(current[0], ithInterval[0])
			current[1] = max(current[1], ithInterval[1])
		} else {
			result = append(result, current)
			current = []int{ithInterval[0], ithInterval[1]}
		}
	}
	result = append(result, current)
	return result
}
