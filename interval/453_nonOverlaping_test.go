package interval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// [[1,2],[2,3],[3,4],[1,3]]
// [[1,2],[1,2],[1,2]]
// [[1,2],[2,3]]

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{
			name:      "first example",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			// intervals: [][]int{{1, 2}, {1, 3}, {2, 3}, {3, 4}},
			expected: 1,
		},
		{
			name:      "second example",
			intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
			expected:  2,
		},
		{
			name:      "third example",
			intervals: [][]int{{1, 2}, {2, 2}},
			expected:  0,
		},
		{
			name:      "fourth example",
			intervals: [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}},
			expected:  2,
		},
		{
			name:      "fourth example",
			intervals: [][]int{{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}},
			expected:  7,
		},
		{
			name: "fifth example",
			intervals: [][]int{
				// {-73, -26, 47}, {-65, -11, 54}, {-63, 2, 65},
				{-62, -49, 13},
				// {-52, 31, 83},
				{-40, -26, 14},
				// {-31, 49, 80},
				{30, 47, 17},
				{58, 95, 37},
				// {66, 98, 32}, {82, 97, 15},
				{95, 99, 4}},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, eraseOverlapIntervals(test.intervals))
		})
	}
}
