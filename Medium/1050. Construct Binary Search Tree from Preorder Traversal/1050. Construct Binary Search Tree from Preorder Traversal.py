# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:

    def bstFromPreorder(self, preorder: List[int]) -> TreeNode:

        root = TreeNode(preorder[0])
        parents = [root]

        for i in range(1, len(preorder)):

            node = TreeNode(preorder[i])

            if node.val < parents[-1].val:
                parents[-1].left = node
                parents.append(node)
                continue

            while len(parents) > 1:
                if node.val < parents[-2].val:
                    break
                parents.pop()

            while parents[-1].right is not None:
                parents.append(parents[-1].right)

            parents[-1].right = node

            parents.append(node)

        return root

