package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "first example",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			name:     "first example",
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
		{
			name:     "third example",
			nums:     []int{-4, -1, 4, -5, 1, -6, 9, -6, 0, 2, 2, 7, 0, 9, -3, 8, 9, -2, -6, 5, 0, 3, 4, -2},
			expected: 12,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, longestConsecutive(test.nums))
		})
	}
}
