# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:

    def __init__(self):

        self.sum = float("-inf")

    def maxPathSumCalc(self, root: TreeNode):

        if root is None:
            return 0

        leftPathSum = self.maxPathSumCalc(root.left)
        rightPathSum = self.maxPathSumCalc(root.right)

        pathSum = max(
            root.val,
            root.val + max(leftPathSum, rightPathSum)
        )

        rootPathSum = max(
            pathSum,
            root.val + leftPathSum + rightPathSum,
        )

        if self.sum < rootPathSum:
            self.sum = rootPathSum

        return pathSum

    def maxPathSum(self, root: TreeNode) -> int:

        self.maxPathSumCalc(root)

        return self.sum
