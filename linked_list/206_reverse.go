package linked_list

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	last := head
	currentHead := head
	head = head.Next
	var aux *ListNode
	for head != nil {
		aux = head.Next
		head.Next = currentHead
		currentHead = head
		head = aux
	}
	last.Next = nil
	return currentHead
}

//
// func reverseList(head *ListNode) *ListNode {
// 	current, prev := head, head
// 	var aux *ListNode
// 	for current != nil {
// 		if current.Next == nil {
// 			prev.Next = aux
// 			break
// 		}
// 		current = current.Next
// 		prev.Next = aux
// 		aux = prev
// 		prev = current
// 	}
// 	return current
// }
