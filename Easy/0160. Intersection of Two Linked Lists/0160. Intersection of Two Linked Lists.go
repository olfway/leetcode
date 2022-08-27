/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    var lenA, lenB int
    pA, pB := headA, headB

    for pA.Next != nil {
        lenA++
        pA = pA.Next
    }

    for pB.Next != nil {
        lenB++
        pB = pB.Next
    }

    if pA != pB {
        return nil
    }

    pA = headA
    for lenA > lenB {
        lenA--
        pA = pA.Next
    }

    pB = headB
    for lenB > lenA {
        lenB--
        pB = pB.Next
    }

    for pA != pB {
        pA = pA.Next
        pB = pB.Next
    }

    return pA
}
