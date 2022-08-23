/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

    var cur, prev *ListNode

    for head != nil {
        cur, head = head, head.Next
        cur.Next, prev = prev, cur
    }

    return cur
}

