package interval

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	result := make([][]int, 0, len(intervals)+1)
	var inserted bool
	var lastInserted []int

	for _, interval := range intervals {
		if inserted {
			result = append(result, interval)
		} else {
			if newInterval[0] > interval[1] {
				result = append(result, interval)
			} else {
				if lastInserted != nil {
					if lastInserted[1] >= interval[0] {
						if interval[1] == max(interval[1], lastInserted[1]) {
							inserted = true
							lastInserted[1] = interval[1]
						}
					} else {
						inserted = true
						result = append(result, interval)
					}
					continue
				}
				if newInterval[0] >= interval[0] {
					lastInserted = []int{interval[0], max(interval[1], newInterval[1])}
					if interval[1] == lastInserted[1] {
						inserted = true
					}
					result = append(result, lastInserted)
				}
			}
		}
	}
	if lastInserted[1] == newInterval[1] {
		inserted = true
	}
	if !inserted {
		result = append(result, newInterval)
	}
	return result
}
