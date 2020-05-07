# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:

    def findValues(self, root, parent=None, depth=0):

        if not root:
            return False

        if depth > self.depth:
            return False

        for value in list(self.values):

            if root.val == value:

                if self.parent:
                    if self.parent != parent:
                        if self.depth == depth:
                            return True
                    return False

                self.depth = depth
                self.parent = parent
                self.values.remove(value)
                break

        if self.findValues(root.left, root, depth + 1):
            return True

        return self.findValues(root.right, root, depth + 1)


    def isCousins(self, root: TreeNode, x: int, y: int) -> bool:

        self.depth = 101
        self.parent = None

        self.values = [x, y]

        return self.findValues(root)

