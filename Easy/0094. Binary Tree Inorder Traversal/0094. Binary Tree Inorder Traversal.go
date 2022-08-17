/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {

    var node *TreeNode
    var stack []*TreeNode
    output :=  make([]int, 0, 0)
    for {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        if len(stack) > 0 {
            node = stack[len(stack)-1]
            root = node.Right
            output = append(output, node.Val)
            stack = stack[:len(stack)-1]
            continue
        }
        return output
    }
}
