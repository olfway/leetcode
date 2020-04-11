# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:

    def height(self, node: TreeNode) -> int:

        if not node:
            return 0

        left = self.height(node.left)
        right = self.height(node.right)

        self.diameter = max(self.diameter, left + right)

        return 1 + max(left, right)


    def diameterOfBinaryTree(self, root: TreeNode) -> int:

        self.diameter = 0

        self.height(root)

        return self.diameter

