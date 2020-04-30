# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:

    def isValid(self, root, arr, index):

        if root is None:
            return False

        if root.val != arr[index]:
            return False

        if index == len(arr) - 1:

            if root.left is not None:
                return False

            if root.right is not None:
                return False

            return True

        if self.isValid(root.left, arr, index + 1):
            return True

        if self.isValid(root.right, arr, index + 1):
            return True

        return False

    def isValidSequence(self, root: TreeNode, arr: List[int]) -> bool:

        return self.isValid(root, arr, 0)

