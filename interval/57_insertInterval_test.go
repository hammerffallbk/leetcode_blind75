package interval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		new       []int
		expected  [][]int
	}{
		{
			name:      "first example",
			intervals: [][]int{{1, 3}, {6, 9}},
			new:       []int{2, 5},
			expected:  [][]int{{1, 5}, {6, 9}},
		},
		{
			name:      "second example",
			intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			new:       []int{4, 8},
			expected:  [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name:      "with empty intervals example",
			intervals: [][]int{},
			new:       []int{4, 8},
			expected:  [][]int{{4, 8}},
		},
		{
			name:      "with two separated intervals",
			intervals: [][]int{{1, 5}},
			new:       []int{6, 8},
			expected:  [][]int{{1, 5}, {6, 8}},
		},
		{
			name:      "contained interval",
			intervals: [][]int{{1, 5}},
			new:       []int{2, 3},
			expected:  [][]int{{1, 5}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, insert(test.intervals, test.new))
		})
	}
}
