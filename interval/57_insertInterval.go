package interval

// insert
// 1.- we verify the edge case when intervals is empty
// 2.- We insert all the intervals that are before our newInterval
// 3.- We reduce all the intervals that are in between
// 4.- We insert all the remaining intervals
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	result := make([][]int, 0, len(intervals)+1)
	var i int
	for i < len(intervals) && newInterval[0] > intervals[i][1] {
		result = append(result, intervals[i])
		i++
	}
	newInt := newInterval
	for i < len(intervals) && intervals[i][0] <= newInt[1] {
		newInt = []int{min(intervals[i][0], newInt[0]), max(intervals[i][1], newInt[1])}
		i++
	}
	result = append(result, newInt)
	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}
	return result
}
