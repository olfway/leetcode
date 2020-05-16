# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:

    def oddEvenList(self, head: ListNode) -> ListNode:

        if head is None:
            return None

        odd_tail = head
        even_head = even_tail = head.next

        if even_head is None:
            return head

        while True:

            even_next = odd_tail.next
            if even_next is not even_tail:
                even_tail.next = even_next
                even_tail = even_next

            if even_next is None:
                break

            odd_next = even_next.next
            if odd_next is None:
                break

            odd_tail.next = odd_next
            odd_tail = odd_next

        odd_tail.next = even_head

        return head

