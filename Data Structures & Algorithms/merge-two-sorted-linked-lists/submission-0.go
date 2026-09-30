/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil && list2 != nil {
		return list2
	} else if list1 != nil && list2 == nil {
		return list1
	}
	// both have values
	var head, curr, curr1, curr2 *ListNode
	
	if list1.Val < list2.Val {
		head = list1
		curr1 = list1.Next
		curr2 = list2
	} else {
		head = list2
		curr2 = list2.Next
		curr1 = list1
	}

	curr = head

	for curr1 != nil && curr2 != nil {
		if curr1.Val < curr2.Val {
			curr.Next = curr1
			curr = curr1
			curr1 = curr1.Next
		} else {
			curr.Next = curr2
			curr = curr2
			curr2 = curr2.Next
		}
	}

	// remainders
	if curr1 != nil {
		for curr1 != nil {
			curr.Next = curr1
			curr = curr1
			curr1 = curr1.Next
		}
	} else if curr2 != nil {
		for curr2 != nil {
			curr.Next = curr2
			curr = curr2
			curr2 = curr2.Next
		}
	}

	return head
}
