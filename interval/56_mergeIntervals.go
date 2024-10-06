package interval

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	result := make([][]int, 0, len(intervals))
	current := []int{intervals[0][0], intervals[0][1]}
	for i := 1; i < len(intervals); i++ {
		ithInterval := []int{intervals[i][0], intervals[i][1]}
		if ithInterval[1] < current[0] {
			current[0], current[1], ithInterval[0], ithInterval[1] = ithInterval[0], ithInterval[1], current[0], current[1]
		}
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
