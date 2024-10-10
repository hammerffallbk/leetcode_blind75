package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		expected []int
	}{
		{
			name:     "first example",
			list:     []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "second example",
			list:     []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "third example",
			list:     []int{},
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, sliceToList(test.expected), reverseList(sliceToList(test.list)))
		})
	}
}

func sliceToList(list []int) *ListNode {
	var head, current *ListNode
	for i := range list {
		if head != nil {
			current.Next = &ListNode{Val: list[i]}
			current = current.Next
		} else {
			head = &ListNode{Val: list[i]}
			current = head
		}
	}
	return head
}
