/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {

    if head.Next == nil {
        return true
    }

    if head.Next.Next == nil {
        return head.Val == head.Next.Val
    }

    var p *ListNode

    p1, p2 := head, head

	for {
		p2 = p2.Next

		if p2.Next == nil {
			p1, p2 = head, p1.Next
			break
		}

		p2 = p2.Next

		if p2.Next == nil {
			p1, p2 = head, p1.Next.Next
			break
		}

		p = head
		head = p1.Next
		p1.Next = p1.Next.Next
		head.Next = p
    }

	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1, p2 = p1.Next, p2.Next
	}

    return true
}

