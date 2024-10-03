package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFinish(t *testing.T) {
	tests := []struct {
		Graph      [][]int
		NumCourses int
		Expected   bool
	}{
		// {
		// 	Graph:      [][]int{{1, 0}},
		// 	NumCourses: 2,
		// 	Expected:   true,
		// },
		{
			Graph:      [][]int{{0, 1}},
			NumCourses: 2,
			Expected:   true,
		},
		// {
		// 	Graph:      [][]int{{1, 0}, {0, 1}},
		// 	NumCourses: 2,
		// 	Expected:   false,
		// },
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, canFinish(test.NumCourses, test.Graph))
	}
}
