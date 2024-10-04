package graph

// longestConsecutive
// for each number on the array we look on the dict if previous and next value
// we update the value on dict, and set init and end of the subsequence
// we take the longest subsequence
func longestConsecutive(nums []int) int {
	dict := make(map[int]int)
	var result int
	for _, num := range nums {
		if _, ok := dict[num]; !ok {
			prev := dict[num-1]
			next := dict[num+1]
			curLen := prev + next + 1
			if prev > 0 {
				dict[num-prev] = curLen
			}
			if next > 0 {
				dict[num+next] = curLen
			}
			dict[num] = curLen
			if curLen > result {
				result = curLen
			}
		}
	}
	return result
}
