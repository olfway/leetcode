/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }
    head, p1, p2 := list1, list1, list2
    if p1.Val > p2.Val {
        head, p1, p2 = list2, list2, list1
    }
    for p2 != nil {
        if p1.Next == nil || p1.Next.Val > p2.Val {
            p2, p1.Next, p2.Next = p2.Next, p2, p1.Next
        }
        p1 = p1.Next
    }
    return head
}
