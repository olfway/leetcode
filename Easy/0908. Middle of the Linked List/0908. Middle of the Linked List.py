# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:

    def middleNode(self, head: ListNode) -> ListNode:

        nodes = [head, ]

        while head.next != None:
            head = head.next
            nodes.append(head)

        return nodes[len(nodes) // 2]

